package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/body"
)

func (h *Handler) getAllDetails(w http.ResponseWriter, r *http.Request) {
	var params body.DetailBodyParams

	query := r.URL.Query()

	name := query.Get("name")

	if name != "" {
		params.Name = name
	}

	minPrice := query.Get("minPrice")

	if minPrice != "" {
		parsedPrice, parseErr := strconv.Atoi(minPrice)

		if parseErr == nil {
			params.MinPrice = parsedPrice
		}
	}

	maxPrice := query.Get("maxPrice")

	if maxPrice != "" {
		parsedPrice, parseErr := strconv.Atoi(maxPrice)

		if parseErr == nil {
			params.MaxPrice = parsedPrice
		}
	}

	detailType := query.Get("type")

	if detailType != "" {
		params.Type = detailType
	}

	details, err := h.detailsService.GetAll(context.Background(), params)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: details,
	})
}

func (h *Handler) createDetail(w http.ResponseWriter, r *http.Request) {
	var detailData body.CreateDetailBody

	err := DecodeRequest(r, &detailData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse Detail data"))
		return
	}

	id, err := h.detailsService.Create(r.Context(), detailData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}

func (h *Handler) deleteDetail(w http.ResponseWriter, r *http.Request) {
	var idBody body.IdBody

	err := DecodeRequest(r, &idBody)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse id"))
		return
	}

	id, err := h.detailsService.Delete(r.Context(), uuid.MustParse(idBody.Id))

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}
