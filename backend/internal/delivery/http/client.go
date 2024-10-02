package handler

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

func (h *Handler) getAllClients(w http.ResponseWriter, r *http.Request) {
	clients, err := h.clientsService.GetAll(context.Background())

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: clients,
	})
}

func (h *Handler) createClient(w http.ResponseWriter, r *http.Request) {
	var clientData entity.CreateClientBody

	err := DecodeRequest(r, &clientData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse client data"))
		return
	}

	err = h.clientsService.Create(r.Context(), clientData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: "ok",
	})
}
