package transport

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"restapi/internal/core"
	"restapi/internal/service"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type httpHandlers struct {
	taskService service.TaskService
	logger      *zap.Logger
}

func NewHTTPHandlers(ts service.TaskService, logger *zap.Logger) *httpHandlers {
	return &httpHandlers{
		taskService: ts,
		logger: logger,
	}
}

func (h *httpHandlers) respond(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error("failed to write http response", zap.Error(err))
	}
}

func (h *httpHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO taskInput
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, io.EOF) {
			errDTO.Message = core.ErrEmptyBody.Error()
		}

		h.respond(w, http.StatusBadRequest, errDTO)
		return
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		h.respond(w, http.StatusBadRequest, errDTO)
		return
	}
	snapshot, err := h.taskService.Create(r.Context(), taskDTO)
	if err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, core.ErrTaskAlreadyExists) {
			h.respond(w, http.StatusConflict, errDTO)
		} else {
			h.respond(w, http.StatusInternalServerError, errDTO)
		}

		return
	}
	h.respond(w, http.StatusCreated, snapshotToResponse(snapshot))
}

func (h *httpHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	title = strings.TrimSpace(title)

	if title == "" {
		errDTO := errorDTO{
			Message: core.ErrEmptyTitle.Error(),
			Time:    time.Now(),
		}
		h.respond(w, http.StatusBadRequest, errDTO)
		return
	}

	snapshot, err := h.taskService.Get(r.Context(), title)
	if err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, core.ErrTaskNotFound) {
			h.respond(w, http.StatusNotFound, errDTO)
		} else {
			h.respond(w, http.StatusInternalServerError, errDTO)
		}
		return
	}
	h.respond(w, http.StatusOK, snapshotToResponse(snapshot))
}

func (h *httpHandlers) HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	completedStr := r.URL.Query().Get("completed")
	var completedFilter *bool
	if completedStr != "" {
		switch completedStr {
		case "true":
			b := true
			completedFilter = &b
		case "false":
			b := false
			completedFilter = &b
		default:
			errDTO := errorDTO{
				Message: core.ErrInvalidCompleted.Error(),
				Time:    time.Now(),
			}
			h.respond(w, http.StatusBadRequest, errDTO)
			return
		}
	}
	snapshots, err := h.taskService.List(r.Context(), completedFilter)
	if err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		h.respond(w, http.StatusInternalServerError, errDTO)
		return
	}
	tasks := make([]taskResponse, 0, len(snapshots))
	for _, s := range snapshots {
		tasks = append(tasks, snapshotToResponse(s))
	}

	h.respond(w, http.StatusOK, tasks)
}

func (h *httpHandlers) HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	var patch patchDTO
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, io.EOF) {
			errDTO.Message = core.ErrEmptyBody.Error()
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	title := mux.Vars(r)["title"]
	title = strings.TrimSpace(title)

	if title == "" {
		errDTO := errorDTO{
			Message: core.ErrEmptyTitle.Error(),
			Time:    time.Now(),
		}
		h.respond(w, http.StatusBadRequest, errDTO)
		return
	}

	snapshot, err := h.taskService.Update(r.Context(), title, patch)
	if err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, core.ErrTaskNotFound) {
			h.respond(w, http.StatusNotFound, errDTO)
		} else {
			h.respond(w, http.StatusInternalServerError, errDTO)
		}

		return
	}

	h.respond(w, http.StatusOK, snapshotToResponse(snapshot))
}

func (h *httpHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	title = strings.TrimSpace(title)

	if title == "" {
		errDTO := errorDTO{
			Message: core.ErrEmptyTitle.Error(),
			Time:    time.Now(),
		}
		h.respond(w, http.StatusBadRequest, errDTO)
		return
	}
	if err := h.taskService.Delete(r.Context(), title); err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, core.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
