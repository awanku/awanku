package coreapi

import (
	"net/http"

	"github.com/awanku/awanku/internal/coreapi/utils/apihelper"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func (s *Server) initRoutes() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("willkomen"))
	})

	s.router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		apihelper.JSON(w, http.StatusOK, s.db.Health())
	})

	s.router.Route("/v1", func(r chi.Router) {
		r.Use(cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodPost, http.MethodPatch, http.MethodDelete},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           5 * 60,
		}).Handler)

		r.Route("/auth", func(r chi.Router) {
			r.Get("/{provider:[a-z]+}/connect", s.authService.HandleGetProviderConnect)
			r.Get("/{provider:[a-z]+}/callback", s.authService.HandleGetProviderCallback)
			r.Post("/token", s.authService.HandlePostToken)
		})

		r.Route("/users", func(r chi.Router) {
			r.Use(s.m.ValidateOauthToken)

			r.Get("/me", s.userService.HandleGetMe)
		})
	})
}