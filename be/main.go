package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/faizallmaullana/rs_gundar/controllers/doctors"
	"github.com/faizallmaullana/rs_gundar/controllers/listdiseases"
	"github.com/faizallmaullana/rs_gundar/controllers/patients"
	"github.com/faizallmaullana/rs_gundar/models"
)

// initilaize the cors middleware
var corsConfig = cors.DefaultConfig()

func init() {
	// allow all origins
	corsConfig.AllowAllOrigins = true
}

func main() {
	r := gin.Default()

	// connect to database
	models.ConnectDatabase()
	r.Use(cors.New(corsConfig))

	// ROUTES

	// ADMIN
	// Administrators

	// PATIENT
	// Patient
	r.GET("/patients", patients.FindPatients)
	r.GET("/patient/:id", patients.FindPatient)
	r.POST("/patient/", patients.CreatePatient)
	r.PATCH("/patient/:id", patients.UpdatePatient)
	r.DELETE("/patient/:id", patients.DeletePatient)
	// recovery
	r.GET("/patients/deleted", patients.FindDeletedPatients)
	r.DELETE("/patient/deleted/:id", patients.RecoverDeletedPatient)

	// DOCTORS
	// Doctor
	r.GET("/doctors", doctors.FindDoctors)
	r.GET("/doctor/:id", doctors.FindDoctor)
	r.POST("/doctor/", doctors.CreateDoctor)
	r.PATCH("/doctor/:id", doctors.UpdateDoctor)
	r.DELETE("/doctor/:id", doctors.DeleteDoctor)
	// recovery
	r.GET("/doctors/deleted", doctors.FindDeletedDoctor)
	r.DELETE("/doctor/deleted/:id", doctors.RecoverDeletedDoctor)

	// LIST OF DISEASE
	// ListOfDiseases
	r.GET("/diseases", listdiseases.FindDiseases)
	r.GET("/disease/:id", listdiseases.FindDisease)
	r.POST("/disease/", listdiseases.CreateDisease)
	r.PATCH("/disease/:id", listdiseases.UpdateDisease)
	r.DELETE("/disease/:id", listdiseases.DeleteDisease)
	// recovery
	r.GET("/diseases/deleted", listdiseases.FindDeletedDiseases)
	r.DELETE("/disease/deleted/:id", listdiseases.RecoverDeletedDisease)

	// run the server
	r.Run()
}
