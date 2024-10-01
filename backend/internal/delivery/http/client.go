package handler

import (
	"net/http"
)

func (h *Handler) getAllClients(w http.ResponseWriter, r *http.Request) {
	clients, err := h.clientsService.GetAllClients()

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: clients,
	})
}
