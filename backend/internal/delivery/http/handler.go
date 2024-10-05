package handler

import (
	"github.com/go-chi/chi"
)

type Handler struct {
	Router              *chi.Mux
	clientsService      ClientsService
	employeesService    EmployeesService
	contractsService    ContractsService
	receiptsService     ReceiptsService
	vehiclesService     VehiclesService
	applicationsService ApplicationsService
}

func New(
	clientsService ClientsService,
	employeesService EmployeesService,
	contractsService ContractsService,
	receiptsService ReceiptsService,
	vehiclesService VehiclesService,
	applicationsService ApplicationsService,
) *Handler {
	h := &Handler{
		Router:           NewRouter(),
		clientsService:   clientsService,
		employeesService: employeesService,
		contractsService: contractsService,
		receiptsService:  receiptsService,
		vehiclesService:  vehiclesService,
	}

	h.Router.Route("/clients", h.ClientSubroutes)
	h.Router.Route("/employees", h.EmployeeSubroutes)
	h.Router.Route("/contracts", h.ContractSubroutes)
	h.Router.Route("/receipts", h.ReceiptSubroutes)
	h.Router.Route("/vehicles", h.VehicleSubroutes)
	h.Router.Route("/applications", h.ApplicationSubroutes)

	return h
}
