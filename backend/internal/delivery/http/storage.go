package handler

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/body"
)

func (h *Handler) getAllStorages(w http.ResponseWriter, r *http.Request) {
	Storage, err := h.storagesService.GetAll(context.Background())

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: Storage,
	})
}

func (h *Handler) createStorage(w http.ResponseWriter, r *http.Request) {
	var StorageData body.CreateStorageBody

	err := DecodeRequest(r, &StorageData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse Storage data"))
		return
	}

	id, err := h.storagesService.Create(r.Context(), StorageData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}
