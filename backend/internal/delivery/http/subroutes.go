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
	r.Patch("/{id}", h.updateEmployee)
}

func (h *Handler) ContractSubroutes(r chi.Router) {
	r.Get("/", h.getAllContracts)
	r.Post("/", h.createContract)
	r.Patch("/{id}", h.updateContract)
}
