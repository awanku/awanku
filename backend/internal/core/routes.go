package core

import (
	"net/http"

	"github.com/awanku/awanku/backend/internal/core/utils/apihelper"
	"github.com/go-chi/chi"
)

func (s *Server) initRoutes() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("willkommen freunden"))
	})

	s.router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		var primaryOK bool
		s.db.WriterQuery(&primaryOK, "select true;")

		var replicaOK bool
		s.db.Query(&replicaOK, "select true;")

		apihelper.JSON(w, http.StatusOK, map[string]interface{}{
			"database": map[string]bool{
				"primary": primaryOK,
				"replica": replicaOK,
			},
		})
	})

	s.router.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Get("/{provider:[a-z]+}/connect", s.authService.HandleGetProviderConnect)
			r.Get("/{provider:[a-z]+}/callback", s.authService.HandleGetProviderCallback)
		})
	})
}
