package handler

import "github.com/go-chi/chi"

func (h *Handler) ClientSubroutes(r chi.Router) {
	r.Get("/", h.getAllClients)
	r.Post("/", h.createClient)
}
