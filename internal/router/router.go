package router

import (
	"fmt"
	"net/http"

	"github.com/findsam/auth-micro/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Handlers struct { 
	User *handler.UserHandler
}

type Router struct {
	addr     string
	handlers *Handlers
}

func New(addr string, h *Handlers) *Router {
	return &Router{
		addr:     fmt.Sprintf(":%s", addr),
		handlers: h,
	}
}

func (s *Router) Start() error {
	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.URLFormat)
	c.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	c.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/sign-up", s.handlers.User.CreateUser)
			r.Route("/user/{id}", func(r chi.Router) {
			})
		})
	})

	return http.ListenAndServe(s.addr, c)
}
