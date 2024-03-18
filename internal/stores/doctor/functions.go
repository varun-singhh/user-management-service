package doctor

import (
	"github.com/varun-singhh/backend/user-management-service/internal/models"
	"strconv"
	"strings"
	"time"
)

func getWhereClause(doctor *models.Doctor) (string, []interface{}) {
	var (
		query  string
		params []interface{}
		i      int
	)

	i = 0

	if doctor.DoctorDetails != nil && doctor.DoctorDetails.Name != "" {
		if i > 0 {
			query += " AND "
		}
		query += "name=$" + strconv.Itoa(i+1)
		params = append(params, doctor.DoctorDetails.Name)
		i++
	}

	if doctor.DoctorDetails != nil && doctor.DoctorDetails.DOB != "" {
		if i > 0 {
			query += " AND "
		}
		query += "dob=$" + strconv.Itoa(i+1)
		params = append(params, doctor.DoctorDetails.DOB)
		i++
	}

	if doctor.DoctorDetails != nil && doctor.DoctorDetails.Age != "" {
		if query != "" {
			query += " AND "
		}
		query += "age=$" + strconv.Itoa(i+1)
		params = append(params, doctor.DoctorDetails.Age)
		i++
	}

	if doctor.DoctorDetails != nil && doctor.DoctorDetails.City != "" {
		if query != "" {
			query += " AND "
		}
		query += "city=$" + strconv.Itoa(i+1)
		params = append(params, doctor.DoctorDetails.City)
		i++
	}

	if doctor.DoctorDetails != nil && doctor.DoctorDetails.State != "" {
		if query != "" {
			query += " AND "
		}
		query += "state=$" + strconv.Itoa(i+1)
		params = append(params, doctor.DoctorDetails.State)
		i++
	}

	if doctor.DoctorDetails != nil && doctor.DoctorDetails.District != "" {
		if query != "" {
			query += " AND "
		}
		query += "district=$" + strconv.Itoa(i+1)
		params = append(params, doctor.DoctorDetails.District)
		i++
	}

	if doctor.DoctorDetails != nil && doctor.DoctorDetails.Pincode != "" {
		if query != "" {
			query += " AND "
		}
		query += "pincode=$" + strconv.Itoa(i+1)
		params = append(params, doctor.DoctorDetails.Pincode)
		i++
	}

	if doctor.DoctorDetails != nil && doctor.DoctorDetails.Gender != "" {
		if query != "" {
			query += " AND "
		}
		query += "gender=$" + strconv.Itoa(i+1)
		params = append(params, doctor.DoctorDetails.Gender)
		i++
	}

	if doctor.DoctorDetails != nil && doctor.DoctorDetails.LicenseNumber != "" {
		if query != "" {
			query += " AND "
		}
		query += "license_number=$" + strconv.Itoa(i+1)
		params = append(params, doctor.DoctorDetails.LicenseNumber)
		i++
	}

	if doctor.DoctorDetails != nil && doctor.DoctorContact.Email != "" {
		if query != "" {
			query += " AND "
		}
		query += "email=$" + strconv.Itoa(i+1)
		params = append(params, doctor.DoctorContact.Email)
		i++
	}

	if doctor.DoctorDetails != nil && doctor.DoctorContact.Phone != "" {
		if query != "" {
			query += " AND "
		}
		query += "phone=$" + strconv.Itoa(i+1)
		params = append(params, doctor.DoctorContact.Phone)
		i++
	}

	if doctor.Department != "" {
		if query != "" {
			query += " AND "
		}
		query += "department=$" + strconv.Itoa(i+1)
		params = append(params, doctor.Department)
		i++
	}

	if doctor.Designation != "" {
		if query != "" {
			query += " AND "
		}
		query += "designation=$" + strconv.Itoa(i+1)
		params = append(params, doctor.Designation)
		i++
	}

	if doctor.ID != "" {
		if query != "" {
			query += " AND "
		}
		query += "id=$" + strconv.Itoa(i+1)
		params = append(params, doctor.ID)
		i++
	}

	return strings.Trim(query, " AND "), params
}

func getInsertColumnsAndPlaceholders(doctor *models.Doctor) (string, string, []interface{}) {
	var (
		columns      []string
		placeholders []string
		params       []interface{}
	)

	if doctor.DoctorDetails != nil {
		if doctor.DoctorDetails.Name != "" {
			columns = append(columns, "name")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, doctor.DoctorDetails.Name)
		}
		if doctor.DoctorDetails.DOB != "" {
			columns = append(columns, "dob")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, doctor.DoctorDetails.DOB)
		}
		if doctor.DoctorDetails.Age != "" {
			columns = append(columns, "age")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, doctor.DoctorDetails.Age)
		}
		if doctor.DoctorDetails.City != "" {
			columns = append(columns, "city")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, doctor.DoctorDetails.City)
		}
		if doctor.DoctorDetails.State != "" {
			columns = append(columns, "state")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, doctor.DoctorDetails.State)
		}
		if doctor.DoctorDetails.District != "" {
			columns = append(columns, "district")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, doctor.DoctorDetails.District)
		}
		if doctor.DoctorDetails.Pincode != "" {
			columns = append(columns, "pincode")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, doctor.DoctorDetails.Pincode)
		}
		if doctor.DoctorDetails.Gender != "" {
			columns = append(columns, "gender")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, doctor.DoctorDetails.Gender)
		}
		if doctor.DoctorDetails.LicenseNumber != "" {
			columns = append(columns, "license_number")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, doctor.DoctorDetails.LicenseNumber)
		}
	}

	if doctor.DoctorContact != nil {
		if doctor.DoctorContact.Email != "" {
			columns = append(columns, "email")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, doctor.DoctorContact.Email)
		}
		if doctor.DoctorContact.Phone != "" {
			columns = append(columns, "phone")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, doctor.DoctorContact.Phone)
		}
	}

	if doctor.ID != "" {
		columns = append(columns, "id")
		placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
		params = append(params, doctor.ID)
	}

	if doctor.Department != "" {
		columns = append(columns, "department")
		placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
		params = append(params, doctor.Department)
	}

	if doctor.Designation != "" {
		columns = append(columns, "designation")
		placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
		params = append(params, doctor.Designation)
	}

	return "(" + strings.Join(columns, ", ") + ")", "(" + strings.Join(placeholders, ", ") + ")", params
}

func getUpdateSetClause(doctor *models.Doctor) (string, []interface{}) {
	var (
		setValues []string
		params    []interface{}
	)

	if doctor.DoctorDetails != nil {
		if doctor.DoctorDetails.Name != "" {
			setValues = append(setValues, "name=$"+strconv.Itoa(len(params)+1))
			params = append(params, doctor.DoctorDetails.Name)
		}
		if doctor.DoctorDetails.DOB != "" {
			setValues = append(setValues, "dob=$"+strconv.Itoa(len(params)+1))
			params = append(params, doctor.DoctorDetails.DOB)
		}
		if doctor.DoctorDetails.Age != "" {
			setValues = append(setValues, "age=$"+strconv.Itoa(len(params)+1))
			params = append(params, doctor.DoctorDetails.Age)
		}
		if doctor.DoctorDetails.City != "" {
			setValues = append(setValues, "city=$"+strconv.Itoa(len(params)+1))
			params = append(params, doctor.DoctorDetails.City)
		}
		if doctor.DoctorDetails.State != "" {
			setValues = append(setValues, "state=$"+strconv.Itoa(len(params)+1))
			params = append(params, doctor.DoctorDetails.State)
		}
		if doctor.DoctorDetails.District != "" {
			setValues = append(setValues, "district=$"+strconv.Itoa(len(params)+1))
			params = append(params, doctor.DoctorDetails.District)
		}
		if doctor.DoctorDetails.Pincode != "" {
			setValues = append(setValues, "pincode=$"+strconv.Itoa(len(params)+1))
			params = append(params, doctor.DoctorDetails.Pincode)
		}
		if doctor.DoctorDetails.Gender != "" {
			setValues = append(setValues, "gender=$"+strconv.Itoa(len(params)+1))
			params = append(params, doctor.DoctorDetails.Gender)
		}
		if doctor.DoctorDetails.LicenseNumber != "" {
			setValues = append(setValues, "license_number=$"+strconv.Itoa(len(params)+1))
			params = append(params, doctor.DoctorDetails.LicenseNumber)
		}
	}

	if doctor.DoctorContact != nil {
		if doctor.DoctorContact.Email != "" {
			setValues = append(setValues, "email=$"+strconv.Itoa(len(params)+1))
			params = append(params, doctor.DoctorContact.Email)
		}
		if doctor.DoctorContact.Phone != "" {
			setValues = append(setValues, "phone=$"+strconv.Itoa(len(params)+1))
			params = append(params, doctor.DoctorContact.Phone)
		}
	}

	if doctor.Department != "" {
		setValues = append(setValues, "department=$"+strconv.Itoa(len(params)+1))
		params = append(params, doctor.Department)
	}

	if doctor.Designation != "" {
		setValues = append(setValues, "designation=$"+strconv.Itoa(len(params)+1))
		params = append(params, doctor.Designation)
	}

	setValues = append(setValues, "updated_at=$"+strconv.Itoa(len(params)+1))
	params = append(params, time.Now())

	setClause := strings.Join(setValues, ", ")
	return setClause, params
}

func getFilterParams(filter *models.DoctorFilter) (query string, params []interface{}) {
	if filter.Name != "" {
		query += "name LIKE $" + strconv.Itoa(len(params)+1)
		params = append(params, "%"+filter.Name+"%")
	}
	if filter.LicenseNumber != "" {
		if query != "" {
			query += " AND "
		}
		query += "license_number LIKE $" + strconv.Itoa(len(params)+1)
		params = append(params, "%"+filter.LicenseNumber+"%")
	}
	if filter.Email != "" {
		if query != "" {
			query += " AND "
		}
		query += "email LIKE $" + strconv.Itoa(len(params)+1)
		params = append(params, "%"+filter.Email+"%")
	}
	if filter.Phone != "" {
		if query != "" {
			query += " AND "
		}
		query += "phone LIKE $" + strconv.Itoa(len(params)+1)
		params = append(params, "%"+filter.Phone+"%")
	}
	return query, params
}
