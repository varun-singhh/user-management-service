package models

import "time"

type Doctor struct {
	ID             string                 `json:"id"`
	DoctorDetails  *DoctorPersonalDetails `json:"doctor_details"`
	DoctorContact  *DoctorContact         `json:"doctor_contact"`
	Department     string                 `json:"department"`
	Designation    string                 `json:"designation"`
	CreatedAt      *time.Time             `json:"created_at"`
	DeletedAt      *time.Time             `json:"deleted_at,omitempty"`
	UpdatedAt      *time.Time             `json:"updated_at,omitempty"`
	FollowUps      []*Followup            `json:"follow_ups"`
	CaseRecordings []*CaseRecording       `json:"case_recordings"`
	Patients       []*Patient             `json:"patients"`
}

type DoctorPersonalDetails struct {
	Name          string `json:"name"`
	DOB           string `json:"DOB"`
	Age           string `json:"age"`
	City          string `json:"city"`
	State         string `json:"state"`
	District      string `json:"district"`
	Pincode       string `json:"pincode"`
	Gender        string `json:"gender"`
	LicenseNumber string `json:"license_number"`
}

type DoctorContact struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type Followup struct {
	ID            string   `json:"id"`
	DoctorID      string   `json:"doctor_id"`
	Date          string   `json:"date,omitempty"`
	Time          string   `json:"time"`
	FollowUpCount string   `json:"followup_count"`
	Patient       *Patient `json:"patient"`
}

type CaseRecording struct {
	ID       string   `json:"id"`
	DoctorID string   `json:"doctor_id"`
	Patient  *Patient `json:"patient"`
}
