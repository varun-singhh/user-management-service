package patient

import (
	"github.com/varun-singhh/backend/user-management-service/internal/models"
)

func validateCreateRequest(patient *models.Patient) []string {
	var params []string

	if patient.PatientDetails.Name == "" {
		params = append(params, "name")
	}

	if patient.PatientDetails.AadharNumber == "" {
		params = append(params, "aadhar_number")
	}

	if patient.PatientDetails.Gender == "" {
		params = append(params, "gender")
	}

	if patient.PatientContact.Phone == "" {
		params = append(params, "phone")
	}

	if patient.PatientContact.Email == "" {
		params = append(params, "email")
	}

	return params
}
