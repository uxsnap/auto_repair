package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/body"
)

func (h *Handler) getAllStorages(w http.ResponseWriter, r *http.Request) {
	var params body.StorageBodyParams

	query := r.URL.Query()

	storageNum := query.Get("storageNum")

	if storageNum != "" {
		params.StorageNum = storageNum
	}

	employeeName := query.Get("employeeName")

	if employeeName != "" {
		params.EmployeeName = employeeName
	}

	detailName := query.Get("detailName")

	if detailName != "" {
		params.DetailName = detailName
	}

	Storage, err := h.storagesService.GetAll(context.Background(), params)

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

func (h *Handler) deleteStorage(w http.ResponseWriter, r *http.Request) {
	var idBody body.IdBody

	err := DecodeRequest(r, &idBody)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse id"))
		return
	}

	id, err := h.storagesService.Delete(r.Context(), uuid.MustParse(idBody.Id))

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}

func (h *Handler) updateStorage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	clientID := uuid.MustParse(id)

	var clientData body.CreateStorageBody

	err := DecodeRequest(r, &clientData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse client data"))
		return
	}

	err = h.storagesService.Update(r.Context(), clientID, clientData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}
