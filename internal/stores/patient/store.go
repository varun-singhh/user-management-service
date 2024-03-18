package patient

import (
	"database/sql"
	"github.com/varun-singhh/backend/user-management-service/internal/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"strconv"
	"time"
)

type storeHandler struct {
}

func New() *storeHandler {
	return &storeHandler{}
}

func (a *storeHandler) Get(ctx *gofr.Context, patientReq *models.Patient) (*models.Patient, error) {
	var patient models.Patient
	query := `SELECT id, email, phone, name, dob, age,gender,relation_name,aadhar_number,city, district, pincode, created_at, updated_at, deleted_at FROM patients `

	// Append the WHERE clause for other conditions
	whereClause, params := getWhereClause(patientReq)
	if whereClause != "" {
		query += "WHERE " + whereClause
	}

	patient.PatientContact = &models.PatientContact{}
	patient.PatientDetails = &models.PatientPersonalDetails{}
	patient.Address = &models.Address{}

	row := ctx.DB().QueryRow(query, params...)
	err := row.Scan(
		&patient.ID,
		&patient.PatientContact.Email,
		&patient.PatientContact.Phone,
		&patient.PatientDetails.Name,
		&patient.PatientDetails.DOB,
		&patient.PatientDetails.Age,
		&patient.PatientDetails.Gender,
		&patient.PatientDetails.Relation,
		&patient.PatientDetails.AadharNumber,
		&patient.Address.City,
		&patient.Address.District,
		&patient.Address.Pincode,
		&patient.CreatedAt,
		&patient.UpdatedAt,
		&patient.DeletedAt,
	)

	if err == sql.ErrNoRows || patient.DeletedAt != nil {
		return nil, &errors.EntityNotFound{
			Entity: "patient",
			ID:     patientReq.ID,
		}
	}

	if err != nil {
		ctx.Logger.Error("error while fetching patient details from db:", err)
		return nil, &errors.DB{Err: err}
	}

	return &patient, nil
}

func (a *storeHandler) GetAll(ctx *gofr.Context, doctorFilter *models.PatientFilter, pageFilter *models.Page) ([]*models.Patient, error) {
	var patients []*models.Patient
	query := `SELECT id, email, phone, name, dob, age,gender,relation_name,aadhar_number,city, district, pincode, created_at, updated_at, deleted_at FROM patients `

	// Append the WHERE clause for other conditions
	whereClause, params := getFilterParams(doctorFilter)
	if whereClause != "" {
		query += "WHERE " + whereClause
	}

	query += " LIMIT " + pageFilter.Limit

	query += " OFFSET " + pageFilter.Offset

	rows, err := ctx.DB().Query(query, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var patient models.Patient
		patient.PatientContact = &models.PatientContact{}
		patient.PatientDetails = &models.PatientPersonalDetails{}
		patient.Address = &models.Address{}

		err = rows.Scan(
			&patient.ID,
			&patient.PatientContact.Email,
			&patient.PatientContact.Phone,
			&patient.PatientDetails.Name,
			&patient.PatientDetails.DOB,
			&patient.PatientDetails.Age,
			&patient.PatientDetails.Gender,
			&patient.PatientDetails.Relation,
			&patient.PatientDetails.AadharNumber,
			&patient.Address.City,
			&patient.Address.District,
			&patient.Address.Pincode,
			&patient.CreatedAt,
			&patient.UpdatedAt,
			&patient.DeletedAt,
		)
		if err != nil {
			ctx.Logger.Error("error while scanning doctor details:", err)
			return nil, &errors.DB{Err: err}
		}

		patients = append(patients, &patient)
	}

	if err = rows.Err(); err != nil {
		ctx.Logger.Error("error while iterating over doctor rows:", err)
		return nil, &errors.DB{Err: err}
	}

	return patients, nil
}

func (a *storeHandler) Create(ctx *gofr.Context, patient *models.Patient) (*models.Patient, error) {
	columns, placeholders, params := getInsertColumnsAndPlaceholders(patient)
	query := "INSERT INTO patients " + columns + " VALUES " + placeholders + " RETURNING id"

	var id string
	err := ctx.DB().QueryRow(query, params...).Scan(&id)
	if err != nil {
		ctx.Logger.Error("error while creating patient in the database:", err)
		return nil, &errors.DB{Err: err}
	}
	patient.ID = id
	return patient, nil
}

func (a *storeHandler) Update(ctx *gofr.Context, patient *models.Patient) (*models.Patient, error) {
	setClause, params := getUpdateSetClause(patient)
	query := "UPDATE patients SET " + setClause + " WHERE id = $" + strconv.Itoa(len(params)+1)
	params = append(params, patient.ID)

	_, err := ctx.DB().Exec(query, params...)
	if err != nil {
		ctx.Logger.Error("error while updating patient in the database:", err)
		return nil, &errors.DB{Err: err}
	}
	return patient, nil
}

func (a *storeHandler) Delete(ctx *gofr.Context, id string) error {
	query := `UPDATE patients  SET deleted_at=$1 WHERE id=$2`
	_, err := ctx.DB().Exec(query, time.Now(), id)
	if err != nil {
		ctx.Logger.Error("error while deleting patient from the database:", err)
		return &errors.DB{Err: err}
	}
	return nil
}

func (a *storeHandler) GetFollowUp(ctx *gofr.Context, patientID string) ([]*models.FollowUp, error) {
	query := ` SELECT id, date, time, followup_count FROM patient_followups  WHERE patient_id=$1`
	rows, err := ctx.DB().Query(query, patientID)
	if err != nil {
		ctx.Logger.Error("error while fetching follow-ups from the database:", err)
		return nil, &errors.DB{Err: err}
	}
	defer rows.Close()

	followUps := []*models.FollowUp{}
	for rows.Next() {
		var followUp models.FollowUp
		err := rows.Scan(&followUp.ID, &followUp.Date, &followUp.Time, &followUp.FollowUpCount)
		if err != nil {
			ctx.Logger.Error("error while scanning follow-up row:", err)
			return nil, &errors.DB{Err: err}
		}
		followUps = append(followUps, &followUp)
	}

	if err := rows.Err(); err != nil {
		ctx.Logger.Error("error while iterating follow-up rows:", err)
		return nil, &errors.DB{Err: err}
	}

	return followUps, nil
}
