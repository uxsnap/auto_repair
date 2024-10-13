package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/body"
)

func (h *Handler) getAllVehicles(w http.ResponseWriter, r *http.Request) {
	var params body.VehicleBodyParams

	query := r.URL.Query()

	vehicleNumber := query.Get("vehicleNumber")

	if vehicleNumber != "" {
		params.VehicleNumber = vehicleNumber
	}

	brand := query.Get("brand")

	if brand != "" {
		params.Brand = brand
	}

	model := query.Get("model")

	if model != "" {
		params.Model = model
	}

	vehicles, err := h.vehiclesService.GetAll(context.Background(), params)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: vehicles,
	})
}

func (h *Handler) createVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicleData body.CreateVehicleBody

	err := DecodeRequest(r, &vehicleData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("недостаточно данных для ТС"))
		return
	}

	id, err := h.vehiclesService.Create(r.Context(), vehicleData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}

func (h *Handler) deleteVehicle(w http.ResponseWriter, r *http.Request) {
	var idBody body.IdBody

	err := DecodeRequest(r, &idBody)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse id"))
		return
	}

	id, err := h.vehiclesService.Delete(r.Context(), uuid.MustParse(idBody.Id))

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}
