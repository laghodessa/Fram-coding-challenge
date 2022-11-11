package http

import (
	"context"
	"database/sql"
	"personia/app"
	"personia/infra/sqlite"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ServerOpts struct {
	DB        *sql.DB
	APISecret string
}

func NewServer(opts ServerOpts) *Server {
	hrUC := app.HRUC{
		HierarchyRepo: sqlite.NewHierarchyRepo(opts.DB),
	}

	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = ErrorHandler()
	e.Use(middleware.Logger())
	e.Use(middleware.KeyAuth(func(auth string, c echo.Context) (bool, error) {
		return auth == opts.APISecret, nil
	}))

	s := &Server{
		e:    e,
		hrUC: hrUC,
	}
	s.registerRoutes()
	return s
}

type Server struct {
	e    *echo.Echo
	hrUC app.HRUC
}

func (s *Server) registerRoutes() {
	api := s.e.Group("/api")
	RegisterHierarchy(api, s.hrUC)
}

func (s *Server) Start(addr string) error {
	return s.e.Start(addr)
}

func (s *Server) Shutdown() error {
	timeout, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return s.e.Shutdown(timeout)
}

func Context(ctx echo.Context) context.Context {
	return ctx.Request().Context()
}
