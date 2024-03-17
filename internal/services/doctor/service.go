package doctor

import (
	"github.com/pkg/errors"
	"github.com/varun-singhh/backend/user-management-service/internal/models"
	"github.com/varun-singhh/backend/user-management-service/internal/stores"
	errors2 "gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"net/http"
)

type serviceHandler struct {
	store stores.Doctor
}

func New(store stores.Doctor) *serviceHandler {
	return &serviceHandler{store: store}
}

func (h *serviceHandler) Get(ctx *gofr.Context, doctor *models.Doctor) (interface{}, error) {
	resp, err := h.store.Get(ctx, doctor)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *serviceHandler) Create(ctx *gofr.Context, doctor *models.Doctor) (interface{}, error) {
	var entityNotFound *errors2.EntityNotFound

	isExistingUser, err := h.store.Get(ctx, &models.Doctor{ID: doctor.ID})
	if !errors.As(err, &entityNotFound) && err != nil {
		return nil, err
	}

	if isExistingUser != nil {
		return nil, &errors2.Response{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: "doctor already exist"}
	}

	resp, err := h.store.Create(ctx, doctor)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *serviceHandler) Update(ctx *gofr.Context, doctor *models.Doctor) (interface{}, error) {
	_, err := h.store.Get(ctx, &models.Doctor{ID: doctor.ID})
	if err != nil {
		return nil, err
	}

	resp, err := h.store.Update(ctx, doctor)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *serviceHandler) Delete(ctx *gofr.Context, id string) error {
	_, err := h.store.Get(ctx, &models.Doctor{ID: id})
	if err != nil {
		return err
	}

	err = h.store.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (h *serviceHandler) GetFollowUp(ctx *gofr.Context, doctorID string) (interface{}, error) {
	resp, err := h.store.GetFollowUp(ctx, doctorID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
