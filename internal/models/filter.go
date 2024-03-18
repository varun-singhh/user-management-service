package models

type DoctorFilter struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	LicenseNumber string `json:"license_number"`
	Email         string `json:"email,omitempty"`
	Phone         string `json:"phone,omitempty"`
}

type PatientFilter struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	AadharNumber string `json:"aadhar_number"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
}
