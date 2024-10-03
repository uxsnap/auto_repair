package handler

import (
	"github.com/go-chi/chi"
)

type Handler struct {
	Router           *chi.Mux
	clientsService   ClientsService
	employeesService EmployeesService
}

func New(
	clientsService ClientsService,
	employeesService EmployeesService,
) *Handler {
	h := &Handler{
		Router:           NewRouter(),
		clientsService:   clientsService,
		employeesService: employeesService,
	}

	h.Router.Route("/clients", h.ClientSubroutes)
	h.Router.Route("/employee", h.EmployeeSubroutes)

	return h
}
