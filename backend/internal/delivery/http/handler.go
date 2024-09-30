package handler

import (
	"github.com/go-chi/chi"
)

type Handler struct {
	Router         *chi.Mux
	clientsService ClientService
}

func New() *Handler {
	h := &Handler{
		Router: NewRouter(),
	}

	h.Router.Route("/client", h.ClientSubroutes)

	return h
}
