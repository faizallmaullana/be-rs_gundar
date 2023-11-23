package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/faizallmaullana/rs_gundar/controllers/admin"
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
	r.GET("/admins", admin.FindAdmins)
	r.GET("/admin/:id", admin.FindAdmin)
	r.POST("/admin/", admin.CreateAdmin)
	r.PATCH("/admin/:id", admin.UpdateAdmin)
	r.DELETE("/admin/:id", admin.DeleteAdmin)
	// recovery
	r.GET("/admins/deleted", admin.FindDeletedAdmins)
	r.DELETE("/admin/deleted/:id", admin.RecoverDeletedAdmin)

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

	// Treatment History
	r.GET("/treatments", patients.FindTreatments)
	r.GET("/treatment/:id", patients.FindTreatment)
	r.POST("/treatment/", patients.CreateTreatment)
	r.PATCH("/treatment/:id", patients.UpdateTreatment)
	r.DELETE("/treatment/:id", patients.DeleteTreatment)
	// recovery
	r.GET("/treatments/deleted", patients.FindDeletedTreatments)
	r.DELETE("/treatment/deleted/:id", patients.RecoverDeletedTreatment)

	// Drug Consumption History
	r.GET("/drugconss", patients.FindDrugconss)
	r.GET("/drugcons/:id", patients.FindDrugcons)
	r.POST("/drugcons/", patients.CreateDrugcons)
	r.PATCH("/drugcons/:id", patients.UpdateDrugcons)
	r.DELETE("/drugcons/:id", patients.DeleteDrugcons)
	// recovery
	r.GET("/drugconss/deleted", patients.FindDeletedDrugconss)
	r.DELETE("/drugcons/deleted/:id", patients.RecoverDeletedDrugcons)

	// HospitalizedHistory
	r.GET("/hospitals", patients.FindHospitals)
	r.GET("/hospital/:id", patients.FindHospital)
	r.POST("/hospital/", patients.CreateHospital)
	r.PATCH("/hospital/:id", patients.UpdateHospital)
	r.DELETE("/hospital/:id", patients.DeleteHospital)
	// recovery
	r.GET("/hospitals/deleted", patients.FindDeletedHospitals)
	r.DELETE("/hospital/deleted/:id", patients.RecoverDeletedHospital)

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

	// Doctor Specialties
	r.GET("/docspecs", doctors.FindDocspecs)
	r.GET("/docspec/:id", doctors.FindDocspec)
	r.POST("/docspec/", doctors.CreateDocspec)
	r.PATCH("/docspec/:id", doctors.UpdateDocspec)
	r.DELETE("/docspec/:id", doctors.DeleteDocspec)
	// recovery
	r.GET("/docspecs/deleted", doctors.FindDeletedDocspecs)
	r.DELETE("/docspec/deleted/:id", doctors.RecoverDeletedDocspec)

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

	// ListOfDrugs
	r.GET("/drugs", listdiseases.FindDrugs)
	r.GET("/drug/:id", listdiseases.FindDrug)
	r.POST("/drug/", listdiseases.CreateDrug)
	r.PATCH("/drug/:id", listdiseases.UpdateDrug)
	r.DELETE("/drug/:id", listdiseases.DeleteDrug)
	// recovery
	r.GET("/drugs/deleted", listdiseases.FindDeletedDrugs)
	r.DELETE("/drug/deleted/:id", listdiseases.RecoverDeletedDrug)

	// ListOfHospitalChambers
	r.GET("/chambers", listdiseases.FindChambers)
	r.GET("/chamber/:id", listdiseases.FindChamber)
	r.POST("/chamber/", listdiseases.CreateChamber)
	r.PATCH("/chamber/:id", listdiseases.UpdateChamber)
	r.DELETE("/chamber/:id", listdiseases.DeleteChamber)
	// recovery
	r.GET("/chambers/deleted", listdiseases.FindDeletedChambers)
	r.DELETE("/chamber/deleted/:id", listdiseases.RecoverDeletedChamber)

	// run the server
	r.Run()
}
