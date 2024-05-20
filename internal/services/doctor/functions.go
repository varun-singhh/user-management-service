package doctor

import (
	"github.com/varun-singhh/backend/user-management-service/internal/models"
)

func validateCreateRequest(doctor *models.Doctor) []string {
	var params []string

	if doctor.DoctorDetails.Name == "" {
		params = append(params, "name")
	}

	if doctor.DoctorDetails.LicenseNumber == "" {
		params = append(params, "license_number")
	}

	if doctor.DoctorDetails.Gender == "" {
		params = append(params, "gender")
	}

	if doctor.DoctorContact.Phone == "" {
		params = append(params, "phone")
	}

	if doctor.DoctorContact.Email == "" {
		params = append(params, "email")
	}

	return params
}
