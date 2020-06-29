package coreapi

import (
	"net/http"
	"os"
	"time"

	hansip "github.com/asasmoyo/pq-hansip"
	"github.com/awanku/awanku/internal/coreapi/auth"
	ourMiddleware "github.com/awanku/awanku/internal/coreapi/middleware"
	"github.com/awanku/awanku/internal/coreapi/user"
	"github.com/awanku/awanku/pkg/core"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v10"
)

type Server struct {
	router      chi.Router
	authService auth.AuthService
	userService user.UserService
    userSettingsService user.UserSettingsService
	m           *ourMiddleware.Middleware
	db          *hansip.Cluster
}

func (s *Server) Init() error {
	s.router = chi.NewRouter()
	s.router.Use(middleware.Logger)

	var oauthTokenSecretKey = []byte("supersecretkey")

	opt, err := pg.ParseURL(os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	s.db = hansip.NewCluster(&hansip.Config{
		Primary:        hansip.WrapGoPG(pg.Connect(opt)),
		Replicas:       []hansip.SQL{hansip.WrapGoPG(pg.Connect(opt))},
		PingTimeout:    1 * time.Second,
		ConnCheckDelay: 5 * time.Second,
	})

	cs, err := core.NewCoreService(&core.Config{
		DB:                  s.db,
		OauthTokenSecretKey: oauthTokenSecretKey,
	})
	if err != nil {
		return err
	}

	s.authService = auth.AuthService{
		OauthTokenSecretKey: oauthTokenSecretKey,
		UserStore:           cs.UserStore(),
		AuthStore:           cs.AuthStore(),
	}
	s.authService.Init()

	s.userService = user.UserService{
		UserStore: cs.UserStore(),
	}
	s.authService.Init()

    s.userSettingsService = user.UserSettingsService{
        UserSettingsStore: cs.UserSettingsStore(),
    }
    s.userSettingsService.Init()

	s.m = &ourMiddleware.Middleware{
		OauthTokenSecretKey: oauthTokenSecretKey,
		AuthStore:           cs.AuthStore(),
	}

	s.initRoutes()
	return nil
}

func (s *Server) Start() error {
	return http.ListenAndServe("0.0.0.0:3000", s.router)
}
