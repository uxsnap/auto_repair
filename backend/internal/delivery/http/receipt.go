package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/uxsnap/auto_repair/backend/internal/body"
)

func (h *Handler) getAllReceipts(w http.ResponseWriter, r *http.Request) {
	Receipts, err := h.receiptsService.GetAll(context.Background())

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: Receipts,
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

func (h *Handler) deleteReceipt(w http.ResponseWriter, r *http.Request) {
	var idBody body.IdBody

	err := DecodeRequest(r, &idBody)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse id"))
		return
	}

	id, err := h.receiptsService.Delete(r.Context(), uuid.MustParse(idBody.Id))

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}

func (h *Handler) updateReceipt(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	ReceiptID := uuid.MustParse(id)

	var ReceiptData body.CreateReceiptBody

	err := DecodeRequest(r, &ReceiptData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("cannot parse Receipt data"))
		return
	}

	err = h.receiptsService.Update(r.Context(), ReceiptID, ReceiptData)

	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	WriteResponseJson(w, DataResponse{
		Data: id,
	})
}
