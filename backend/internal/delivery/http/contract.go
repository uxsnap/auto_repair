package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/body"
)

func (h *Handler) getAllContracts(w http.ResponseWriter, r *http.Request) {
	var params body.ContractBodyParams

	query := r.URL.Query()

	name := query.Get("name")

	if name != "" {
		params.Name = name
	}

	minSum := query.Get("minSum")

	if minSum != "" {
		parsedSum, parseErr := strconv.Atoi(minSum)

		if parseErr == nil {
			params.MinSum = parsedSum
		}
	}

	maxSum := query.Get("maxSum")

	if maxSum != "" {
		parsedSum, parseErr := strconv.Atoi(maxSum)

		if parseErr == nil {
			params.MaxSum = parsedSum
		}
	}

	status := query.Get("status")

	if status != "" {
		params.Status = status
	}

	minCreatedAt := query.Get("minCreatedAt")

	if minCreatedAt != "" {
		params.MinCreatedAt = minCreatedAt
	}

	maxCreatedAt := query.Get("maxCreatedAt")

	if maxCreatedAt != "" {
		params.MaxCreatedAt = maxCreatedAt
	}

	contracts, err := h.contractsService.GetAll(context.Background(), params)

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
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("невозможно распарсить тело запроса"))
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
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("невозможно распарсить тело запроса"))
		return
	}

	fmt.Println(clientID, contractData)

	err = h.contractsService.Update(r.Context(), clientID, contractData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}

func (h *Handler) deleteContract(w http.ResponseWriter, r *http.Request) {
	var idBody body.IdBody

	err := DecodeRequest(r, &idBody)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("невозможно распарсить тело запроса"))
		return
	}

	id, err := h.contractsService.Delete(r.Context(), uuid.MustParse(idBody.Id))

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}
