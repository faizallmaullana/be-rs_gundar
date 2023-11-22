package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/faizallmaullana/rs_gundar/controllers"
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
	// Patient
	r.GET("/patients", controllers.FindPatients)
	r.GET("/patient/:id", controllers.FindPatient)
	r.POST("/patient/", controllers.CreatePatient)
	r.PATCH("/patient/:id", controllers.UpdatePatient)
	r.DELETE("/patient/:id", controllers.DeletePatient)
	// recovery
	r.GET("/patients/deleted", controllers.FindDeletedPatients)
	r.DELETE("/patient/deleted/:id", controllers.RecoverDeletedPatient)

	// run the server
	r.Run()
}
