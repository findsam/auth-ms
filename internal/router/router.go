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
	User  *handler.UserHandler
	Store *handler.StoreHandler
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
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	c.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/sign-up", s.handlers.User.SignUp)
			r.Post("/sign-in", s.handlers.User.SignIn)
			r.Route("/user", func(r chi.Router) {
				r.Get("/{id:[a-f0-9]{24}}", s.handlers.User.GetById)
				r.Get("/{username:[^a-f0-9].*}", s.handlers.User.GetByUsername)
			})
			r.Group(func(r chi.Router) {
				r.Use(WithJWT)
				r.Get("/me", s.handlers.User.Me)
				r.Get("/refresh", s.handlers.User.Refresh)
			})
		})

		r.Route("/stores", func(r chi.Router) {
			r.Group(func(r chi.Router){
				r.Use(WithJWT)
				r.Post("/", s.handlers.Store.Create)
			})
		})
	})

	return http.ListenAndServe(s.addr, c)
}
