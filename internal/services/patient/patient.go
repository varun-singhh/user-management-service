package patient

import (
	"errors"
	"github.com/varun-singhh/backend/user-management-service/internal/models"
	"github.com/varun-singhh/backend/user-management-service/internal/stores"
	errors2 "gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"net/http"
)

type serviceHandler struct {
	store stores.Patient
}

func New(store stores.Patient) *serviceHandler {
	return &serviceHandler{store: store}
}

func (h *serviceHandler) Get(ctx *gofr.Context, patient *models.Patient) (interface{}, error) {
	resp, err := h.store.Get(ctx, patient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *serviceHandler) Create(ctx *gofr.Context, patient *models.Patient) (interface{}, error) {
	var entityNotFound *errors2.EntityNotFound

	isExistingUser, err := h.store.Get(ctx, &models.Patient{ID: patient.ID})
	if !errors.As(err, &entityNotFound) && err != nil {
		return nil, err
	}

	if isExistingUser != nil {
		return nil, &errors2.Response{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: "patient already exist"}
	}

	resp, err := h.store.Create(ctx, patient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *serviceHandler) Update(ctx *gofr.Context, patient *models.Patient) (interface{}, error) {
	_, err := h.store.Get(ctx, &models.Patient{ID: patient.ID})
	if err != nil {
		return nil, err
	}

	resp, err := h.store.Update(ctx, patient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *serviceHandler) Delete(ctx *gofr.Context, id string) error {
	_, err := h.store.Get(ctx, &models.Patient{ID: id})
	if err != nil {
		return err
	}

	err = h.store.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (h *serviceHandler) GetFollowUp(ctx *gofr.Context, patientID string) (interface{}, error) {
	resp, err := h.store.GetFollowUp(ctx, patientID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
