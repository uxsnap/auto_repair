package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/body"
)

func (h *Handler) getAllReceipts(w http.ResponseWriter, r *http.Request) {
	var params body.ReceiptBodyParams

	query := r.URL.Query()

	contractName := query.Get("contractName")

	if contractName != "" {
		params.ContractName = contractName
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

	receipts, err := h.receiptsService.GetAll(context.Background(), params)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: receipts,
	})
}

func (h *Handler) createReceipt(w http.ResponseWriter, r *http.Request) {
	var ReceiptData body.CreateReceiptBody

	err := DecodeRequest(r, &ReceiptData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse Receipt data"))
		return
	}

	id, err := h.receiptsService.Create(r.Context(), ReceiptData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}
