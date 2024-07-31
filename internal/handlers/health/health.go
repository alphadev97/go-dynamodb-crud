package health

import (
	"errors"
	"net/http"

	"github.com/alphadev97/go-dynamodb-crud/internal/handlers"
	"github.com/alphadev97/go-dynamodb-crud/internal/repository/adapter"
)

type Handler struct {
	handlers.Interface
	Respository adapter.Interface
}

func NewHandler(respository adapter.Interface) handlers.Interface {
	return &Handler{
		Respository: respository,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if !h.Respository.Health() {
		HttpStatus.StatusIntervalServerError(w, r, errors.New("Relational database not alive"))
		return
	}

	HttpStatus.StatusOK(w, r, "Service OK")
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}
