package patient

import (
	"github.com/varun-singhh/backend/user-management-service/internal/models"
	"strconv"
	"strings"
	"time"
)

func getWhereClause(patient *models.Patient) (string, []interface{}) {
	var (
		query  string
		params []interface{}
		i      int
	)

	i = 0

	if patient.PatientDetails != nil {
		if patient.PatientDetails.Name != "" {
			if i > 0 {
				query += " AND "
			}
			query += "name=$" + strconv.Itoa(i+1)
			params = append(params, patient.PatientDetails.Name)
			i++
		}
		if patient.PatientDetails.DOB != "" {
			if i > 0 {
				query += " AND "
			}
			query += "dob=$" + strconv.Itoa(i+1)
			params = append(params, patient.PatientDetails.DOB)
			i++
		}
		if patient.PatientDetails.Age != "" {
			if query != "" {
				query += " AND "
			}
			query += "age=$" + strconv.Itoa(i+1)
			params = append(params, patient.PatientDetails.Age)
			i++
		}
		if patient.PatientDetails.Gender != "" {
			if query != "" {
				query += " AND "
			}
			query += "gender=$" + strconv.Itoa(i+1)
			params = append(params, patient.PatientDetails.Gender)
			i++
		}
		if patient.PatientDetails.AadharNumber != "" {
			if query != "" {
				query += " AND "
			}
			query += "aadhar_number=$" + strconv.Itoa(i+1)
			params = append(params, patient.PatientDetails.AadharNumber)
			i++
		}
	}

	if patient.Address != nil {
		if patient.Address.Block != "" {
			if query != "" {
				query += " AND "
			}
			query += "block=$" + strconv.Itoa(i+1)
			params = append(params, patient.Address.Block)
			i++
		}

		if patient.Address.State != "" {
			if query != "" {
				query += " AND "
			}
			query += "state=$" + strconv.Itoa(i+1)
			params = append(params, patient.Address.State)
			i++
		}
		if patient.Address.District != "" {
			if query != "" {
				query += " AND "
			}
			query += "district=$" + strconv.Itoa(i+1)
			params = append(params, patient.Address.District)
			i++
		}
		if patient.Address.Pincode != "" {
			if query != "" {
				query += " AND "
			}
			query += "pincode=$" + strconv.Itoa(i+1)
			params = append(params, patient.Address.Pincode)
			i++
		}
	}

	if patient.PatientContact != nil {
		if patient.PatientContact.Email != "" {
			if query != "" {
				query += " AND "
			}
			query += "email=$" + strconv.Itoa(i+1)
			params = append(params, patient.PatientContact.Email)
			i++
		}
		if patient.PatientContact.Phone != "" {
			if query != "" {
				query += " AND "
			}
			query += "phone=$" + strconv.Itoa(i+1)
			params = append(params, patient.PatientContact.Phone)
			i++
		}
	}

	if patient.ID != "" {
		if query != "" {
			query += " AND "
		}
		query += "id=$" + strconv.Itoa(i+1)
		params = append(params, patient.ID)
		i++
	}

	return strings.Trim(query, " AND "), params
}

func getInsertColumnsAndPlaceholders(patient *models.Patient) (string, string, []interface{}) {
	var (
		columns      []string
		placeholders []string
		params       []interface{}
	)

	if patient.PatientDetails != nil {
		if patient.PatientDetails.Name != "" {
			columns = append(columns, "name")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, patient.PatientDetails.Name)
		}
		if patient.PatientDetails.DOB != "" {
			columns = append(columns, "dob")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, patient.PatientDetails.DOB)
		}
		if patient.PatientDetails.Age != "" {
			columns = append(columns, "age")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, patient.PatientDetails.Age)
		}
		if patient.PatientDetails.Gender != "" {
			columns = append(columns, "gender")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, patient.PatientDetails.Gender)
		}
		if patient.PatientDetails.AadharNumber != "" {
			columns = append(columns, "aadhar_number")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, patient.PatientDetails.AadharNumber)
		}
		if patient.PatientDetails.Relation != "" {
			columns = append(columns, "relation_name")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, patient.PatientDetails.Relation)
		}
	}

	if patient.Address != nil {
		if patient.Address.Block != "" {
			columns = append(columns, "block")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, patient.Address.Block)
		}
		if patient.Address.State != "" {
			columns = append(columns, "state")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, patient.Address.State)
		}
		if patient.Address.District != "" {
			columns = append(columns, "district")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, patient.Address.District)
		}
		if patient.Address.Pincode != "" {
			columns = append(columns, "pincode")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, patient.Address.Pincode)
		}
	}

	if patient.PatientContact != nil {
		if patient.PatientContact.Email != "" {
			columns = append(columns, "email")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, patient.PatientContact.Email)
		}

		if patient.PatientContact.Phone != "" {
			columns = append(columns, "phone")
			placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
			params = append(params, patient.PatientContact.Phone)
		}
	}

	if patient.ID != "" {
		columns = append(columns, "id")
		placeholders = append(placeholders, "$"+strconv.Itoa(len(placeholders)+1))
		params = append(params, patient.ID)
	}

	return "(" + strings.Join(columns, ", ") + ")", "(" + strings.Join(placeholders, ", ") + ")", params
}

func getUpdateSetClause(patient *models.Patient) (string, []interface{}) {
	var (
		setValues []string
		params    []interface{}
	)

	if patient.PatientDetails != nil {
		if patient.PatientDetails.Name != "" {
			setValues = append(setValues, "name=$"+strconv.Itoa(len(params)+1))
			params = append(params, patient.PatientDetails.Name)
		}
		if patient.PatientDetails.DOB != "" {
			setValues = append(setValues, "dob=$"+strconv.Itoa(len(params)+1))
			params = append(params, patient.PatientDetails.DOB)
		}
		if patient.PatientDetails.Age != "" {
			setValues = append(setValues, "age=$"+strconv.Itoa(len(params)+1))
			params = append(params, patient.PatientDetails.Age)
		}
		if patient.PatientDetails.Gender != "" {
			setValues = append(setValues, "gender=$"+strconv.Itoa(len(params)+1))
			params = append(params, patient.PatientDetails.Gender)
		}
		if patient.PatientDetails.AadharNumber != "" {
			setValues = append(setValues, "aadhar_number=$"+strconv.Itoa(len(params)+1))
			params = append(params, patient.PatientDetails.AadharNumber)
		}
		if patient.PatientDetails.Relation != "" {
			setValues = append(setValues, "relation_name=$"+strconv.Itoa(len(params)+1))
			params = append(params, patient.PatientDetails.Relation)
		}
	}

	if patient.Address != nil {
		if patient.Address.Block != "" {
			setValues = append(setValues, "block=$"+strconv.Itoa(len(params)+1))
			params = append(params, patient.Address.Block)
		}
		if patient.Address.State != "" {
			setValues = append(setValues, "state=$"+strconv.Itoa(len(params)+1))
			params = append(params, patient.Address.State)
		}
		if patient.Address.District != "" {
			setValues = append(setValues, "district=$"+strconv.Itoa(len(params)+1))
			params = append(params, patient.Address.District)
		}
		if patient.Address.Pincode != "" {
			setValues = append(setValues, "pincode=$"+strconv.Itoa(len(params)+1))
			params = append(params, patient.Address.Pincode)
		}
	}

	if patient.PatientContact != nil {
		if patient.PatientContact.Email != "" {
			setValues = append(setValues, "email=$"+strconv.Itoa(len(params)+1))
			params = append(params, patient.PatientContact.Email)
		}
		if patient.PatientContact.Phone != "" {
			setValues = append(setValues, "phone=$"+strconv.Itoa(len(params)+1))
			params = append(params, patient.PatientContact.Phone)
		}
	}

	setValues = append(setValues, "updated_at=$"+strconv.Itoa(len(params)+1))
	params = append(params, time.Now())

	setClause := strings.Join(setValues, ", ")
	return setClause, params
}

func getFilterParams(filter *models.PatientFilter) (query string, params []interface{}) {
	if filter.Name != "" {
		query += "name LIKE $" + strconv.Itoa(len(params)+1)
		params = append(params, "%"+filter.Name+"%")
	}
	if filter.AadharNumber != "" {
		if query != "" {
			query += " AND "
		}
		query += "aadhar_number LIKE $" + strconv.Itoa(len(params)+1)
		params = append(params, "%"+filter.AadharNumber+"%")
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
