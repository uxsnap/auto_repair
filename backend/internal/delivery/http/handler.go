package handler

import (
	"github.com/go-chi/chi"
)

type Handler struct {
	Router           *chi.Mux
	clientsService   ClientsService
	employeesService EmployeesService
	contractsService ContractsService
}

func New(
	clientsService ClientsService,
	employeesService EmployeesService,
	contractsService ContractsService,
) *Handler {
	h := &Handler{
		Router:           NewRouter(),
		clientsService:   clientsService,
		employeesService: employeesService,
		contractsService: contractsService,
	}

	h.Router.Route("/clients", h.ClientSubroutes)
	h.Router.Route("/employees", h.EmployeeSubroutes)
	h.Router.Route("/contracts", h.ContractSubroutes)

	return h
}
