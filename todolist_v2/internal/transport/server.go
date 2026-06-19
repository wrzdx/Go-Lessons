package transport

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
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
	if err := http.ListenAndServe(":8000", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
