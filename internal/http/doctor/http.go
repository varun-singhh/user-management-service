package doctor

import (
	"github.com/varun-singhh/backend/user-management-service/internal/models"
	"github.com/varun-singhh/backend/user-management-service/internal/services"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	service services.Doctor
}

func New(svc services.Doctor) *handler {
	return &handler{service: svc}
}

func (h *handler) Get(ctx *gofr.Context) (interface{}, error) {
	var doctor models.Doctor

	params := ctx.Request().URL.Query()

	doctor.ID = ctx.PathParam("id")

	doctor.DoctorDetails.Name = params.Get("name")
	doctor.DoctorDetails.LicenseNumber = params.Get("license_number")
	doctor.DoctorContact.Phone = params.Get("phone")

	resp, err := h.service.Get(ctx, &doctor)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	var (
		doctor models.DoctorFilter
		page   models.Page
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

	doctor.Name = params.Get("name")
	doctor.LicenseNumber = params.Get("license_number")
	doctor.Phone = params.Get("phone")

	resp, err := h.service.GetAll(ctx, &doctor, &page)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *handler) Create(ctx *gofr.Context) (interface{}, error) {
	var doctor models.Doctor

	doctor.ID = ctx.PathParam("id")

	err := ctx.Bind(&doctor)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.Create(ctx, &doctor)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *handler) InternalCreate(ctx *gofr.Context) (interface{}, error) {
	var doctor models.Doctor

	err := ctx.Bind(&doctor)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.InternalCreate(ctx, &doctor)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *handler) Update(ctx *gofr.Context) (interface{}, error) {
	var doctor models.Doctor

	doctor.ID = ctx.PathParam("id")

	err := ctx.Bind(&doctor)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.Update(ctx, &doctor)
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
