package models

import "time"

type Patient struct {
	ID             string                  `json:"id"`
	PatientDetails *PatientPersonalDetails `json:"patient_details"`
	PatientContact *PatientContact         `json:"patient_contact"`
	Address        *Address                `json:"address"`
	CreatedAt      *time.Time              `json:"created_at"`
	DeletedAt      *time.Time              `json:"deleted_at,omitempty"`
	UpdatedAt      *time.Time              `json:"updated_at,omitempty"`
	FollowUps      []*FollowUp             `json:"follow_ups"`
	CaseRecordings []*PCaseRecording       `json:"case_recordings"`
	Doctors        []*Doctor               `json:"doctors"`
}

type PatientPersonalDetails struct {
	Name         string `json:"name"`
	DOB          string `json:"DOB"`
	Age          string `json:"age"`
	Gender       string `json:"gender"`
	AadharNumber string `json:"aadhar_number"`
	Relation     string `json:"relation_name"`
}

type PatientContact struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type Address struct {
	District string `json:"district"`
	State    string `json:"state"`
	Block    string `json:"block,omitempty"`
	Pincode  string `json:"pincode"`
}

type FollowUp struct {
	ID            string  `json:"id"`
	Date          string  `json:"date,omitempty"`
	Time          string  `json:"time"`
	FollowUpCount string  `json:"follow_up_count"`
	Doctor        *Doctor `json:"doctor"`
}

type PCaseRecording struct {
	ID     string  `json:"id"`
	Doctor *Doctor `json:"doctor"`
}
