package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/body"
)

func (h *Handler) getAllContracts(w http.ResponseWriter, r *http.Request) {
	contracts, err := h.contractsService.GetAll(context.Background())

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: contracts,
	})
}

func (h *Handler) createContract(w http.ResponseWriter, r *http.Request) {
	var contractData body.CreateContractBody

	err := DecodeRequest(r, &contractData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse client data"))
		return
	}

	id, err := h.contractsService.Create(r.Context(), contractData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}

func (h *Handler) updateContract(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	clientID := uuid.MustParse(id)

	var contractData body.CreateContractBody

	err := DecodeRequest(r, &contractData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse client data"))
		return
	}

	err = h.contractsService.Update(r.Context(), clientID, contractData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}
