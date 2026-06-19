package transport

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type httpServer struct {
	httpHandlers *httpHandlers
}

func NewHTTPServer(httpHandler *httpHandlers) *httpServer {
	return &httpServer{
		httpHandlers: httpHandler,
	}
}

func (s *httpServer) StartServer() error {
	router := mux.NewRouter()
	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTask)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTasks)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleUpdateTask)
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteTask)
	router.Use(s.httpHandlers.LoggingMiddleware)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		s.httpHandlers.logger.Info("HTTP server is running on :8000")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.httpHandlers.logger.Fatal("could not listen on :8000", zap.Error(err))
		}
	}()
	sig := <-quit
	s.httpHandlers.logger.Info("Shutdown signal received", zap.String("signal", sig.String()))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.httpHandlers.logger.Info("Shutting down server gracefully...")
	if err := srv.Shutdown(ctx); err != nil {
		s.httpHandlers.logger.Error("Server forced to shutdown", zap.Error(err))
		return err
	}
	s.httpHandlers.logger.Info("Server exited properly")
	return nil
}
