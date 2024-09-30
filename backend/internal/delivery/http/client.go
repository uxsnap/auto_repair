package handler

import (
	"context"
	"net/http"
)

func (h *Handler) getAllClients(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	h.clientsService.GetAllClients(ctx)
}
