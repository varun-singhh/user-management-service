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

	params := ctx.Request().URL.Query()

	patient.PatientDetails.Name = params.Get("name")
	patient.PatientDetails.AadharNumber = params.Get("aadhar_number")
	patient.PatientContact.Phone = params.Get("phone")

	resp, err := h.service.Get(ctx, &patient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	var (
		patient models.PatientFilter
		page    models.Page
	)

	params := ctx.Request().URL.Query()

	if params.Get("limit") != "" {
		page.Limit = params.Get("limit")
	} else {
		page.Limit = "20"
	}

	if params.Get("offset") != "" {
		page.Offset = params.Get("offset")
	} else {
		page.Offset = "0"
	}

	patient.Name = params.Get("name")
	patient.AadharNumber = params.Get("aadhar_number")
	patient.Phone = params.Get("phone")

	resp, err := h.service.GetAll(ctx, &patient, &page)
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

func (h *handler) InternalCreate(ctx *gofr.Context) (interface{}, error) {
	var patient models.Patient

	err := ctx.Bind(&patient)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.InternalCreate(ctx, &patient)
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
