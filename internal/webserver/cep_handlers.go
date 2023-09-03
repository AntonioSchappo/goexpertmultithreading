package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AntonioSchappo/desafiomultithreading/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ConcurrentRequester interface {
	GetCep(cep string) (entity.Message, error)
}

type CepHandler struct {
	requester ConcurrentRequester
}

func NewCepHandler(r ConcurrentRequester) *CepHandler {
	return &CepHandler{requester: r}
}

func (h *CepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	message, err := h.requester.GetCep(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}
