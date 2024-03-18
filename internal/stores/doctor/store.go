package doctor

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

func (a *storeHandler) Get(ctx *gofr.Context, doctorReq *models.Doctor) (*models.Doctor, error) {
	var doctor models.Doctor
	query := `SELECT id, email, phone, department, designation, name, dob, age, city, district, pincode,license_number, created_at, updated_at, deleted_at FROM doctors `

	// Append the WHERE clause for other conditions
	whereClause, params := getWhereClause(doctorReq)
	if whereClause != "" {
		query += " WHERE " + whereClause
	}

	doctor.DoctorContact = &models.DoctorContact{}
	doctor.DoctorDetails = &models.DoctorPersonalDetails{}

	row := ctx.DB().QueryRow(query, params...)
	err := row.Scan(
		&doctor.ID,
		&doctor.DoctorContact.Email,
		&doctor.DoctorContact.Phone,
		&doctor.Department,
		&doctor.Designation,
		&doctor.DoctorDetails.Name,
		&doctor.DoctorDetails.DOB,
		&doctor.DoctorDetails.Age,
		&doctor.DoctorDetails.City,
		&doctor.DoctorDetails.District,
		&doctor.DoctorDetails.Pincode,
		&doctor.DoctorDetails.LicenseNumber,
		&doctor.CreatedAt,
		&doctor.UpdatedAt,
		&doctor.DeletedAt,
	)

	if err == sql.ErrNoRows || doctor.DeletedAt != nil {
		return nil, &errors.EntityNotFound{
			Entity: "doctor",
			ID:     doctorReq.ID,
		}
	}

	if err != nil {
		ctx.Logger.Error("error while fetching doctor details from db:", err)
		return nil, &errors.DB{Err: err}
	}

	return &doctor, nil
}

func (a *storeHandler) GetAll(ctx *gofr.Context, doctorFilter *models.DoctorFilter, pageFilter *models.Page) ([]*models.Doctor, error) {
	var doctors []*models.Doctor
	query := `SELECT id, email, phone, department, designation, name, dob, age, city, district, pincode, license_number, created_at, updated_at, deleted_at FROM doctors `

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
		var doctor models.Doctor
		doctor.DoctorContact = &models.DoctorContact{}
		doctor.DoctorDetails = &models.DoctorPersonalDetails{}

		err = rows.Scan(
			&doctor.ID,
			&doctor.DoctorContact.Email,
			&doctor.DoctorContact.Phone,
			&doctor.Department,
			&doctor.Designation,
			&doctor.DoctorDetails.Name,
			&doctor.DoctorDetails.DOB,
			&doctor.DoctorDetails.Age,
			&doctor.DoctorDetails.City,
			&doctor.DoctorDetails.District,
			&doctor.DoctorDetails.Pincode,
			&doctor.DoctorDetails.LicenseNumber,
			&doctor.CreatedAt,
			&doctor.UpdatedAt,
			&doctor.DeletedAt,
		)
		if err != nil {
			ctx.Logger.Error("error while scanning doctor details:", err)
			return nil, &errors.DB{Err: err}
		}

		doctors = append(doctors, &doctor)
	}

	if err = rows.Err(); err != nil {
		ctx.Logger.Error("error while iterating over doctor rows:", err)
		return nil, &errors.DB{Err: err}
	}

	return doctors, nil
}

func (a *storeHandler) Create(ctx *gofr.Context, doctor *models.Doctor) (*models.Doctor, error) {
	columns, placeholders, params := getInsertColumnsAndPlaceholders(doctor)
	query := "INSERT INTO doctors " + columns + " VALUES " + placeholders + " RETURNING id"

	var id string
	err := ctx.DB().QueryRow(query, params...).Scan(&id)
	if err != nil {
		ctx.Logger.Error("error while creating doctor in the database:", err)
		return nil, &errors.DB{Err: err}
	}
	doctor.ID = id
	return doctor, nil
}

func (a *storeHandler) Update(ctx *gofr.Context, doctor *models.Doctor) (*models.Doctor, error) {
	setClause, params := getUpdateSetClause(doctor)
	query := "UPDATE doctors SET " + setClause + " WHERE id = $" + strconv.Itoa(len(params)+1)
	params = append(params, doctor.ID)

	_, err := ctx.DB().Exec(query, params...)
	if err != nil {
		ctx.Logger.Error("error while updating doctor in the database:", err)
		return nil, &errors.DB{Err: err}
	}
	return doctor, nil
}

func (a *storeHandler) Delete(ctx *gofr.Context, id string) error {
	query := `UPDATE doctors SET deleted_at=$1 WHERE id=$2`
	_, err := ctx.DB().Exec(query, time.Now(), id)
	if err != nil {
		ctx.Logger.Error("error while deleting doctor from the database:", err)
		return &errors.DB{Err: err}
	}
	return nil
}

func (a *storeHandler) GetFollowUp(ctx *gofr.Context, doctorID string) ([]*models.FollowUp, error) {
	query := ` SELECT id, date, time, followup_count FROM doctor_followups  WHERE doctor_id=$1`
	rows, err := ctx.DB().Query(query, doctorID)
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
