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
	router.Path("/start").Methods("POST").HandlerFunc(s.handlers.HandleStartCompany)
	router.Path("/finish").Methods("POST").HandlerFunc(s.handlers.HandleFinishCompany)

	router.Path("/staff").Methods("GET").Queries("active", "true").HandlerFunc(s.handlers.HandleGetActiveStaff)
	router.Path("/staff").Methods("GET").HandlerFunc(s.handlers.HandleGetAllStaff)
	router.Path("/balance").Methods("GET").HandlerFunc(s.handlers.HandleGetBalance)
	router.Path("/equipments").Methods("GET").HandlerFunc(s.handlers.HandleGetEquipments)
	router.Path("/miners-info").Methods("GET").HandlerFunc(s.handlers.HandleGetMinersInfo)
	router.Path("/stats").Methods("GET").HandlerFunc(s.handlers.HandleGetStats)

	router.Path("/buy").Methods("POST").HandlerFunc(s.handlers.HandleBuyEquipment)
	router.Path("/hire").Methods("POST").HandlerFunc(s.handlers.HandleHireMiner)

	if err := http.ListenAndServe(":8000", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
