package coreapi

import (
	"net/http"

	"github.com/awanku/awanku/internal/coreapi/appctx"
	"github.com/awanku/awanku/internal/coreapi/auth"
	"github.com/awanku/awanku/internal/coreapi/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func (s *Server) initRoutes() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("See https://api.awanku.id/docs/ for API documentation"))
	})

	s.router.Get("/status", statusHandler(s.db))

	s.router.Route("/v1", func(r chi.Router) {
		r.Use(appctx.Middleware(appctx.Config{
			Environment: s.Config.Environment,
			DB:          s.db,
		}))

		r.Use(cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodPost, http.MethodPatch, http.MethodDelete},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           5 * 60,
		}).Handler)

		r.Route("/auth", func(r chi.Router) {
			r.Get("/{provider:[a-z]+}/connect", auth.HandleOauthProviderConnect)
			r.Get("/{provider:[a-z]+}/callback", auth.HandleOauthProviderCallback)
			r.Post("/token", auth.HandleExchangeOauthToken(s.oauthTokenSecretKey))
		})

		r.Route("/users", func(r chi.Router) {
			r.Use(auth.OauthTokenValidatorMiddleware(s.oauthTokenSecretKey))

			r.Get("/me", user.HandleGetMe)
		})
	})
}
