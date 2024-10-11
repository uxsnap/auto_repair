package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/body"
)

func (h *Handler) getAllEmployees(w http.ResponseWriter, r *http.Request) {
	Employees, err := h.employeesService.GetAll(context.Background())

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: Employees,
	})
}

func (h *Handler) createEmployee(w http.ResponseWriter, r *http.Request) {
	var EmployeeData body.CreateEmployeeBody

	err := DecodeRequest(r, &EmployeeData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse Employee data"))
		return
	}

	id, err := h.employeesService.Create(r.Context(), EmployeeData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}

func (h *Handler) updateEmployee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	EmployeeID := uuid.MustParse(id)

	var EmployeeData body.CreateEmployeeBody

	err := DecodeRequest(r, &EmployeeData)

	fmt.Println(EmployeeData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse Employee data"))
		return
	}

	err = h.employeesService.Update(r.Context(), EmployeeID, EmployeeData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}

func (h *Handler) deleteEmployee(w http.ResponseWriter, r *http.Request) {
	var idBody body.IdBody

	err := DecodeRequest(r, &idBody)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse id"))
		return
	}

	id, err := h.employeesService.Delete(r.Context(), uuid.MustParse(idBody.Id))

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}
