package handler

import (
	"github.com/go-chi/chi"
)

func (h *Handler) ClientSubroutes(r chi.Router) {
	r.Get("/", h.getAllClients)
	r.Post("/", h.createClient)
	r.Delete("/", h.deleteClient)
	r.Patch("/{id}", h.updateClient)
}

func (h *Handler) EmployeeSubroutes(r chi.Router) {
	r.Get("/", h.getAllEmployees)
	r.Post("/", h.createEmployee)
	r.Delete("/", h.deleteEmployee)
	r.Patch("/{id}", h.updateEmployee)
}

func (h *Handler) ContractSubroutes(r chi.Router) {
	r.Get("/", h.getAllContracts)
	r.Post("/", h.createContract)
	r.Patch("/{id}", h.updateContract)
}

func (h *Handler) ReceiptSubroutes(r chi.Router) {
	r.Get("/", h.getAllReceipts)
	r.Post("/", h.createReceipt)
	r.Delete("/", h.deleteReceipt)
	r.Patch("/{id}", h.updateReceipt)
}

func (h *Handler) VehicleSubroutes(r chi.Router) {
	r.Get("/", h.getAllVehicles)
	r.Post("/", h.createVehicle)
	r.Delete("/", h.deleteVehicle)
}

func (h *Handler) ApplicationSubroutes(r chi.Router) {
	r.Get("/", h.getAllApplications)
	r.Post("/", h.createApplication)
	r.Delete("/", h.deleteApplication)
}

func (h *Handler) DetailsSubroutes(r chi.Router) {
	r.Get("/", h.getAllDetails)
	r.Post("/", h.createDetail)
	r.Delete("/", h.deleteDetail)
}

func (h *Handler) ActSubroutes(r chi.Router) {
	r.Get("/", h.getAllActs)
	r.Post("/", h.createAct)
	r.Delete("/", h.deleteAct)
}

func (h *Handler) StoragesSubroutes(r chi.Router) {
	r.Get("/", h.getAllStorages)
	r.Post("/", h.createStorage)
	r.Delete("/", h.deleteStorage)
	r.Patch("/{id}", h.updateStorage)
}

func (h *Handler) ServicesSubroutes(r chi.Router) {
	r.Get("/", h.getAllServices)
}
