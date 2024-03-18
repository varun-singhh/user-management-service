package services

import (
	"github.com/varun-singhh/backend/user-management-service/internal/models"
	"gofr.dev/pkg/gofr"
)

type Doctor interface {
	Get(ctx *gofr.Context, doctor *models.Doctor) (interface{}, error)
	GetAll(ctx *gofr.Context, filter *models.DoctorFilter, pageFilter *models.Page) (interface{}, error)
	Create(ctx *gofr.Context, doctor *models.Doctor) (interface{}, error)
	InternalCreate(ctx *gofr.Context, doctor *models.Doctor) (interface{}, error)
	Update(ctx *gofr.Context, doctor *models.Doctor) (interface{}, error)
	Delete(ctx *gofr.Context, id string) error
	GetFollowUp(ctx *gofr.Context, doctorID string) (interface{}, error)
}

type Patient interface {
	Get(ctx *gofr.Context, patient *models.Patient) (interface{}, error)
	GetAll(ctx *gofr.Context, filter *models.PatientFilter, pageFilter *models.Page) (interface{}, error)
	Create(ctx *gofr.Context, patient *models.Patient) (interface{}, error)
	Update(ctx *gofr.Context, patient *models.Patient) (interface{}, error)
	Delete(ctx *gofr.Context, id string) error
	GetFollowUp(ctx *gofr.Context, patientID string) (interface{}, error)
}
