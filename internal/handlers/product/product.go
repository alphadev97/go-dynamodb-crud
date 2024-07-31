package product

import (
	"errors"
	"net/http"

	"github.com/alphadev97/go-dynamodb-crud/internal/handlers"
	"github.com/alphadev97/go-dynamodb-crud/internal/handlers/product"
	"github.com/alphadev97/go-dynamodb-crud/internal/repository/adapter"
)

type Handler struct {
	handlers.Interface
	Controller product.Interface
	Rules      Rules.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: product.NewController(respository),
		Rules:      RulesProduct.NewRules(),
	}
}

func Get()

func getOne()

func getAll()

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	productBody, err := h.getBodyAndValidate(r, uuid.Nil)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}

	ID, err := h.Controller.Create(productBody)

	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, map[string]interface{}{"id": ID.String()})
}

func Put()

func Delete()

func Options()

func (h *Handler) getBodyAndValidate() {}
