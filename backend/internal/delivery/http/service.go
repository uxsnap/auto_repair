package handler

import (
	"context"
	"net/http"
)

func (h *Handler) getAllServices(w http.ResponseWriter, r *http.Request) {
	services, err := h.servicesService.GetAll(context.Background())

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: services,
	})
}
