package server

import (
	"net/http"

	"github.com/thedasken/goth_template/internal/handlers"
	"github.com/thedasken/goth_template/internal/views"
	"github.com/thedasken/goth_template/web"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", handlers.HealthHandler(s.db))

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)
	r.Get("/", templ.Handler(views.HelloForm()).ServeHTTP)
	r.Post("/hello", handlers.HelloWebHandler)

	return r
}
