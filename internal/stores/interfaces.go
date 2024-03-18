package stores

import (
	"github.com/varun-singhh/backend/user-management-service/internal/models"
	"gofr.dev/pkg/gofr"
)

type Doctor interface {
	Get(ctx *gofr.Context, doctorReq *models.Doctor) (*models.Doctor, error)
	GetAll(ctx *gofr.Context, doctorFilter *models.DoctorFilter, pageFilter *models.Page) ([]*models.Doctor, error)
	Create(ctx *gofr.Context, doctor *models.Doctor) (*models.Doctor, error)
	Update(ctx *gofr.Context, doctor *models.Doctor) (*models.Doctor, error)
	Delete(ctx *gofr.Context, id string) error
	GetFollowUp(ctx *gofr.Context, doctorID string) ([]*models.FollowUp, error)
}

type Patient interface {
	Get(ctx *gofr.Context, patientReq *models.Patient) (*models.Patient, error)
	GetAll(ctx *gofr.Context, doctorFilter *models.PatientFilter, pageFilter *models.Page) ([]*models.Patient, error)
	Create(ctx *gofr.Context, patient *models.Patient) (*models.Patient, error)
	Update(ctx *gofr.Context, patient *models.Patient) (*models.Patient, error)
	Delete(ctx *gofr.Context, id string) error
	GetFollowUp(ctx *gofr.Context, patientID string) ([]*models.FollowUp, error)
}
