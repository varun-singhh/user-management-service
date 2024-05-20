package main

import (
	"gofr.dev/pkg/gofr"

	doctorHandler "github.com/varun-singhh/backend/user-management-service/internal/http/doctor"
	doctorSvc "github.com/varun-singhh/backend/user-management-service/internal/services/doctor"
	doctorStore "github.com/varun-singhh/backend/user-management-service/internal/stores/doctor"

	patientHandler "github.com/varun-singhh/backend/user-management-service/internal/http/patient"
	patientSvc "github.com/varun-singhh/backend/user-management-service/internal/services/patient"
	patientStore "github.com/varun-singhh/backend/user-management-service/internal/stores/patient"
)

func main() {
	app := gofr.New()

	ds := doctorStore.New()
	ps := patientStore.New()

	dSvc := doctorSvc.New(ds)
	pSvc := patientSvc.New(ps)

	dh := doctorHandler.New(dSvc)
	ph := patientHandler.New(pSvc)

	app.GET("/doctor/{id}", dh.Get)
	app.GET("/doctors", dh.GetAll)
	app.POST("/doctor/{id}", dh.Create)
	app.POST("/internal/doctor", dh.InternalCreate)
	app.PUT("/doctor/{id}", dh.Update)
	app.DELETE("/doctor/{id}", dh.Delete)
	app.GET("/doctor/followup/{id}", dh.GetFollowUp)

	app.GET("/patient/{id}", ph.Get)
	app.GET("/patients", ph.GetAll)
	app.POST("/patient/{id}", ph.Create)
	app.POST("/internal/patient", ph.InternalCreate)
	app.PUT("/patient/{id}", ph.Update)
	app.DELETE("/patient/{id}", ph.Delete)
	app.GET("/patient/followup/{id}", ph.GetFollowUp)

	app.Start()
}
