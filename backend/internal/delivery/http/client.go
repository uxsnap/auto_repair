package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/body"
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
	var clientData body.CreateClientBody

	err := DecodeRequest(r, &clientData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse client data"))
		return
	}

	id, err := h.clientsService.Create(r.Context(), clientData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}

func (h *Handler) deleteClient(w http.ResponseWriter, r *http.Request) {
	var idBody body.IdBody

	err := DecodeRequest(r, &idBody)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse id"))
		return
	}

	id, err := h.clientsService.Delete(r.Context(), uuid.MustParse(idBody.Id))

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}

func (h *Handler) updateClient(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	clientID := uuid.MustParse(id)

	var clientData body.CreateClientBody

	err := DecodeRequest(r, &clientData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse client data"))
		return
	}

	err = h.clientsService.Update(r.Context(), clientID, clientData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}
