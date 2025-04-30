package api

import (
	"github.com/Deepjyoti-Sarmah/sol-kit-backend/api/handlers"
	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		// r.Get("/health", s.handleHealthCheck)
		r.Get("/health", handlers.HealthCheck)

		// TODO: auth routes
	})

	r.Group(func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
		})

		r.Route("/goals", func(r chi.Router) {
		})

		r.Route("/tasks", func(r chi.Router) {
		})

		r.Route("/funds", func(r chi.Router) {
		})

		r.Route("/challenges", func(r chi.Router) {
		})
	})
}
