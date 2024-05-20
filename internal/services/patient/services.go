package patient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/varun-singhh/backend/user-management-service/internal/models"
	"github.com/varun-singhh/backend/user-management-service/internal/stores"
	errors2 "gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"io/ioutil"
	"net/http"
	"strconv"
)

type serviceHandler struct {
	store stores.Patient
}

func New(store stores.Patient) *serviceHandler {
	return &serviceHandler{store: store}
}

func (h *serviceHandler) Get(ctx *gofr.Context, patient *models.Patient) (interface{}, error) {
	resp, err := h.store.Get(ctx, patient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *serviceHandler) GetAll(ctx *gofr.Context, filter *models.PatientFilter, pageFilter *models.Page) (interface{}, error) {
	resp, err := h.store.GetAll(ctx, filter, pageFilter)
	if err != nil {
		return nil, err
	}

	return models.AllDataResponse{
		Count:  strconv.Itoa(len(resp)),
		Offset: pageFilter.Offset,
		Limit:  pageFilter.Limit,
		Data:   resp,
	}, nil
}

func (h *serviceHandler) Create(ctx *gofr.Context, patient *models.Patient) (interface{}, error) {
	var entityNotFound *errors2.EntityNotFound

	isExistingUser, err := h.store.Get(ctx, &models.Patient{ID: patient.ID})
	if !errors.As(err, &entityNotFound) && err != nil {
		return nil, err
	}

	if isExistingUser != nil {
		return nil, &errors2.Response{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: "patient already exist"}
	}

	resp, err := h.store.Create(ctx, patient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *serviceHandler) InternalCreate(ctx *gofr.Context, patient *models.Patient) (interface{}, error) {
	var entityNotFound *errors2.EntityNotFound

	missingParams := validateCreateRequest(patient)
	if len(missingParams) > 0 {
		return nil, errors2.MissingParam{Param: missingParams}
	}

	// Prepare the request body
	reg := models.Register{
		Phone:      patient.PatientContact.Phone,
		Email:      patient.PatientContact.Email,
		Password:   "welcome@2024",
		Permission: "PATIENT",
	}

	reqBytes, err := json.Marshal(reg)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling: %w", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ctx.Config.Get("AUTHENTICATION_API_ENDPOINT"), bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("error while creating request for auth-service: %w", err)
	}

	// Perform the HTTP request
	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Check if the response status code is not 201 (Created)
	if res.StatusCode != http.StatusCreated {
		var errBody struct {
			Errors []*errors2.Response `json:"errors"`
		}
		err = json.Unmarshal(body, &errBody)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling error body: %w", err)
		}
		return nil, errBody.Errors[0]
	}

	// Unmarshal response data
	var data models.Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	// Check if the user already exists
	isExistingUser, err := h.store.Get(ctx, &models.Patient{ID: strconv.Itoa(data.Data.User.ID)})
	if err != nil && !errors.As(err, &entityNotFound) {
		return nil, fmt.Errorf("error checking if doctor exists: %w", err)
	}

	if isExistingUser != nil {
		return nil, &errors2.Response{StatusCode: http.StatusBadRequest, Code: http.StatusText(http.StatusBadRequest), Reason: "doctor already exists"}
	}

	// Set doctor ID
	patient.ID = strconv.Itoa(data.Data.User.ID)

	// Create the doctor
	resp, err := h.store.Create(ctx, patient)
	if err != nil {
		return nil, fmt.Errorf("error creating doctor: %w", err)
	}

	return resp, nil
}

func (h *serviceHandler) Update(ctx *gofr.Context, patient *models.Patient) (interface{}, error) {
	_, err := h.store.Get(ctx, &models.Patient{ID: patient.ID})
	if err != nil {
		return nil, err
	}

	resp, err := h.store.Update(ctx, patient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *serviceHandler) Delete(ctx *gofr.Context, id string) error {
	_, err := h.store.Get(ctx, &models.Patient{ID: id})
	if err != nil {
		return err
	}

	err = h.store.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (h *serviceHandler) GetFollowUp(ctx *gofr.Context, patientID string) (interface{}, error) {
	resp, err := h.store.GetFollowUp(ctx, patientID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
