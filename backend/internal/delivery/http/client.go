package handler

import (
	"context"
	"net/http"
)

func (h *Handler) getAllClients(w http.ResponseWriter, r *http.Request) {
	clients, err := h.clientsService.GetAllClients(context.Background())

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: clients,
	})
}
