package core

import (
	"net/http"
	"os"
	"time"

	hansip "github.com/asasmoyo/pq-hansip"
	"github.com/awanku/awanku/backend/internal/core/auth"
	"github.com/awanku/awanku/backend/pkg/datastore"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v10"
)

type Server struct {
	router      chi.Router
	authService auth.Auth
	db          *hansip.Cluster
}

func (s *Server) Init() error {
	s.router = chi.NewRouter()
	s.router.Use(middleware.Logger)

	opt, err := pg.ParseURL(os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	s.db = &hansip.Cluster{
		Primary:        hansip.WrapGoPG(pg.Connect(opt)),
		Replicas:       []hansip.SQL{hansip.WrapGoPG(pg.Connect(opt))},
		PingTimeout:    1 * time.Second,
		ConnCheckDelay: 5 * time.Second,
	}
	s.db.Init()

	ds := datastore.DataStore{DB: s.db}
	ds.Init()

	s.authService = auth.Auth{
		Storage: ds.UserStore(),
	}
	s.authService.Init()

	s.initRoutes()
	return nil
}

func (s *Server) Start() error {
	return http.ListenAndServe("0.0.0.0:3000", s.router)
}
