package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/body"
)

func (h *Handler) getAllActs(w http.ResponseWriter, r *http.Request) {
	var params body.ActBodyParams

	query := r.URL.Query()

	name := query.Get("name")

	if name != "" {
		params.Name = name
	}

	applicationName := query.Get("applicationName")

	if applicationName != "" {
		params.ApplicationName = applicationName
	}

	serviceName := query.Get("serviceName")

	if serviceName != "" {
		params.ServiceName = serviceName
	}

	phone := query.Get("phone")

	if phone != "" {
		params.ServiceName = phone
	}

	minCreatedAt := query.Get("minCreatedAt")

	if minCreatedAt != "" {
		params.MinCreatedAt = minCreatedAt
	}

	maxCreatedAt := query.Get("maxCreatedAt")

	if maxCreatedAt != "" {
		params.MaxCreatedAt = maxCreatedAt
	}

	Act, err := h.actsService.GetAll(context.Background(), params)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: Act,
	})
}

func (h *Handler) createAct(w http.ResponseWriter, r *http.Request) {
	var ActData body.CreateActBody

	err := DecodeRequest(r, &ActData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse Act data"))
		return
	}

	id, err := h.actsService.Create(r.Context(), ActData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}

func (h *Handler) deleteAct(w http.ResponseWriter, r *http.Request) {
	var idBody body.IdBody

	err := DecodeRequest(r, &idBody)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse id"))
		return
	}

	id, err := h.actsService.Delete(r.Context(), uuid.MustParse(idBody.Id))

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}
