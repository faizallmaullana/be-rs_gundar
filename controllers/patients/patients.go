package patients

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/rs_gundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePatientInput struct {
	ID        string `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthdate"`
	Address   string `json:"address"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type UpdatePatientInput struct {
	ID        string `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthdate"`
	Address   string `json:"address"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

// PATIENTS

// GET /patients
// get all patients
func FindPatients(c *gin.Context) {
	var dt []models.Patient
	models.DB.Where("is_deleted = ?", false).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// GET /patient/:id
// get a patient by id
func FindPatient(c *gin.Context) {
	// get model if its exists
	var dt models.Patient
	if err := models.DB.Where("id = ? AND is_deleted = ?", c.Param("id"), false).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})

}

// POST /patient
// Create new patient
func CreatePatient(c *gin.Context) {
	var input CreatePatientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Auto-generate ID using UUID
	input.ID = uuid.New().String()

	// Auto-generate CreatedAt to UTC+7
	input.CreatedAt = time.Now().UTC().Add(7 * time.Hour)

	// Auto-set IsDeleted as 0
	input.IsDeleted = false

	layout := "01-02-2006"
	parsedTime, err := time.Parse(layout, input.BirthDate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Birthdate goes wrong"})
		return
	}

	// Create patient
	dt := models.Patient{
		ID:        input.ID,
		Name:      input.Name,
		Gender:    input.Gender,
		BirthDate: parsedTime,
		Address:   input.Address,
		CreatedAt: input.CreatedAt,
		IsDeleted: input.IsDeleted,
	}

	models.DB.Create(&dt)

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// PATCH /patient/:id
// update patient
func UpdatePatient(c *gin.Context) {
	// get model if exist
	var dt models.Patient
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdatePatientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	layout := "01-02-2006"
	parsedTime, err := time.Parse(layout, input.BirthDate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Birthdate goes wrong"})
		return
	}

	data := models.Patient{
		Name:      input.Name,
		Gender:    input.Gender,
		BirthDate: parsedTime,
		Address:   input.Address,
	}

	models.DB.Model(&dt).Updates(data)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /patient/delete/:id
// delete a patient, set IsDeleted as true
func DeletePatient(c *gin.Context) {
	var dt models.Patient
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(UpdatePatientInput{IsDeleted: true})

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}

// RECOVERY

// GET /patient/deleted
// get deleted patient
func FindDeletedPatients(c *gin.Context) {
	var dt []models.Patient
	models.DB.Where("is_deleted = ?", true).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /patient/deleted/:id
// recover deleted patient
func RecoverDeletedPatient(c *gin.Context) {
	var dt models.Patient
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_deleted": false})

	c.JSON(http.StatusOK, gin.H{"data": dt})
}
