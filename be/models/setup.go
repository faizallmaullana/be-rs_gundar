package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "medicalRecord.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	// admin
	database.AutoMigrate(
		&Administrator{},
	)

	// doctor
	database.AutoMigrate(
		&Doctor{},
		&ListOfSpecialties{},
	)

	// patient
	database.AutoMigrate(
		&Patient{},
		&TreatmentHistory{},
		&DrugConsumptionHistory{},
		&HospitalizedHistory{},
	)

	// listDiseases
	database.AutoMigrate(
		&ListOfDiseases{},
		&ListOfDrugs{},
		&ListOfHospitalChambers{},
	)

	DB = database
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
