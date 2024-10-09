package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/body"
)

func (h *Handler) getAllApplications(w http.ResponseWriter, r *http.Request) {
	applications, err := h.applicationsService.GetAll(context.Background())

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: applications,
	})
}

func (h *Handler) createApplication(w http.ResponseWriter, r *http.Request) {
	var ApplicationData body.CreateApplicationBody

	err := DecodeRequest(r, &ApplicationData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse Application data"))
		return
	}

	id, err := h.applicationsService.Create(r.Context(), ApplicationData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}

func (h *Handler) deleteApplication(w http.ResponseWriter, r *http.Request) {
	var idBody body.IdBody

	err := DecodeRequest(r, &idBody)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse id"))
		return
	}

	id, err := h.applicationsService.Delete(r.Context(), uuid.MustParse(idBody.Id))

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}
