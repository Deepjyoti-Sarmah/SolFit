package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Deepjyoti-Sarmah/sol-kit-backend/api/handlers"
	"github.com/Deepjyoti-Sarmah/sol-kit-backend/config"
	db "github.com/Deepjyoti-Sarmah/sol-kit-backend/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	router   *chi.Mux
	config   *config.Config
	db       *pgxpool.Pool
	queries  *db.Querier
	shutdown chan os.Signal
}

func NewServer(cfg *config.Config, pool *pgxpool.Pool, queries *db.Querier) *Server {
	s := &Server{
		router:   chi.NewRouter(),
		config:   cfg,
		db:       pool,
		queries:  queries,
		shutdown: make(chan os.Signal, 1),
	}

	s.setupMiddlewares()

	s.setupRoutes()

	return s
}

func (s *Server) setupMiddlewares() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.RequestID)

	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}

func (s *Server) setupRoutes() {
	// s.router.Get("/health", s.handleHealthCheck)
	s.router.Get("/health", handlers.HealthCheck)

	s.router.Route("/api/v1", func(r chi.Router) {
		// TODO: add routes
	})
}

// func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"status":"ok","message":"server is healthy"}`))
// }

func (s *Server) Start() error {
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.config.ApiServerHost, s.config.ApiServerPort),
		Handler:      s.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("API server listening on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	signal.Notify(s.shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case <-s.shutdown:
		log.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
