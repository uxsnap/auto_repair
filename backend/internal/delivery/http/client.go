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
	ctx := r.Context()
	clientData, ok := ctx.Value("client").(*entity.CreateClientBody)

	if !ok {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse client data"))
		return
	}

	err := h.clientsService.Create(ctx, clientData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: "ok",
	})
}
