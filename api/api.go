package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/HotPotatoC/my-unsplash/internal/json"
	"github.com/HotPotatoC/my-unsplash/internal/logger"

	"github.com/HotPotatoC/my-unsplash/backend"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	ctx  context.Context
	srv  *http.Server
	quit chan os.Signal
}

func NewServer(ctx context.Context, backend backend.Backend) *Server {
	mux := chi.NewMux()

	handler := newHandler(backend)

	mux.Get("/photos", handler.ListPhotos(ctx))
	mux.Get("/photos/search", handler.SearchPhotos(ctx))
	mux.Post("/photos", handler.PostPhoto(ctx))
	mux.Delete("/photos/{id}", handler.DeletePhoto(ctx))

	srvParams := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: mux,
	}

	return &Server{
		ctx:  ctx,
		srv:  srvParams,
		quit: make(chan os.Signal, 1),
	}
}

func (s *Server) Serve() {
	signal.Notify(s.quit, os.Interrupt)

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	log.Printf("Listening on %s", s.srv.Addr)

	<-s.quit

	if err := s.srv.Shutdown(s.ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped")
}

type handler struct {
	backend backend.Backend
}

func newHandler(backend backend.Backend) handler {
	return handler{backend: backend}
}

type JSON map[string]interface{}

func ReplyJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	json, err := json.Marshal(v)
	if err != nil {
		logger.S().Error("failed marshalling json: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}
