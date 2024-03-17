package patient

import (
	"github.com/varun-singhh/backend/user-management-service/internal/models"
	"github.com/varun-singhh/backend/user-management-service/internal/services"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	service services.Patient
}

func New(svc services.Patient) *handler {
	return &handler{service: svc}
}

func (h *handler) Get(ctx *gofr.Context) (interface{}, error) {
	var patient models.Patient

	patient.ID = ctx.PathParam("id")

	err := ctx.Bind(&patient)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.Get(ctx, &patient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *handler) Create(ctx *gofr.Context) (interface{}, error) {
	var patient models.Patient

	patient.ID = ctx.PathParam("id")

	err := ctx.Bind(&patient)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.Create(ctx, &patient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *handler) Update(ctx *gofr.Context) (interface{}, error) {
	var patient models.Patient

	patient.ID = ctx.PathParam("id")

	err := ctx.Bind(&patient)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.Update(ctx, &patient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *handler) Delete(ctx *gofr.Context) (interface{}, error) {
	err := h.service.Delete(ctx, ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *handler) GetFollowUp(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.service.GetFollowUp(ctx, ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
