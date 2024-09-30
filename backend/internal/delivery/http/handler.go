package handler

import (
	"github.com/go-chi/chi"
)

type Handler struct {
	Router         *chi.Mux
	clientsService ClientsService
}

func New(
	clientsService ClientsService,
) *Handler {
	h := &Handler{
		Router:         NewRouter(),
		clientsService: clientsService,
	}

	h.Router.Route("/clients", h.ClientSubroutes)

	return h
}
