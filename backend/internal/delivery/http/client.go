package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) getAllClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	clients, err := h.clientsService.GetAllClients()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	items, marshalErr := json.Marshal(clients)

	if marshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(items)
}
