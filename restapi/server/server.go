package server

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	handlers *HTTPHandlers
}

func NewServer(handlers *HTTPHandlers) *Server {
	return &Server{handlers: handlers}
}

func (s *Server) Start() error {
	router := mux.NewRouter()
	router.Path("/books").Methods("POST").HandlerFunc(s.handlers.HandleCreateBook)
	router.Path("/books/{title}").Methods("GET").HandlerFunc(s.handlers.HandleGetBook)
	router.Path("/books").Methods("GET").Queries("read", "false").HandlerFunc(s.handlers.HandleGetAllUnreadBooks)

	router.Path("/books").Methods("GET").HandlerFunc(s.handlers.HandleGetAllBooks)
	router.Path("/books/{title}").Methods("PATCH").HandlerFunc(s.handlers.HandleReadBook)
	router.Path("/books/{title}").Methods("DELETE").HandlerFunc(s.handlers.HandleDeleteBook)
	
	if err := http.ListenAndServe(":8000", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
