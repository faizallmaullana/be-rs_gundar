package patients

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/rs_gundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateHospitalizedInput struct {
	ID string `json:"id" gorm:"primary_key"`

	// foreign
	PatientID        string `json:"patient_id"`
	Patient          string `json:"patient" gorm:"foreignKey:PatientID"`
	DoctorID         string `json:"doctor_id"`
	Doctor           string `json:"doctor" gorm:"foreignKey:DoctorID"`
	ListOfDiseasesID string `json:"list_of_diseases_id"`
	ListOfDiseases   string `json:"list_of_diseases" gorm:"foreignKey:ListOfDiseasesID"`

	// status
	IsOut     bool      `json:"is_out"`
	CreatedAt time.Time `json:"created_at"`
	OutAt     time.Time `json:"out_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type UpdateHospitalizedInput struct {
	ID string `json:"id" gorm:"primary_key"`

	// foreign
	PatientID        string `json:"patient_id"`
	Patient          string `json:"patient" gorm:"foreignKey:PatientID"`
	DoctorID         string `json:"doctor_id"`
	Doctor           string `json:"doctor" gorm:"foreignKey:DoctorID"`
	ListOfDiseasesID string `json:"list_of_diseases_id"`
	ListOfDiseases   string `json:"list_of_diseases" gorm:"foreignKey:ListOfDiseasesID"`

	// status
	IsOut     bool      `json:"is_out"`
	CreatedAt time.Time `json:"created_at"`
	OutAt     time.Time `json:"out_at"`
	IsDeleted bool      `json:"is_deleted"`
}

// HOSTPITALIZED HISTORY

// GET /hospitals
// get all hospitals
func FindHospitals(c *gin.Context) {
	var dt []models.HospitalizedHistory
	if err := models.DB.Preload("Patient").Preload("Doctor").Preload("ListOfDiseases").
		Where("is_deleted = ?", false).
		Find(&dt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve hospitals"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// GET /hospital/:id
// get a hospital by id
func FindHospital(c *gin.Context) {
	// get model if it exists
	var dt models.HospitalizedHistory
	if err := models.DB.Preload("Patient").Preload("Doctor").Preload("ListOfDiseases").
		Where("id = ? AND is_deleted = ?", c.Param("id"), false).
		First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// POST /hospital
// Create new hospital
func CreateHospital(c *gin.Context) {
	var input CreateHospitalizedInput
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

	// IsOut
	input.IsOut = false

	// Create hospital
	dt := models.HospitalizedHistory{
		ID:               input.ID,
		PatientID:        input.PatientID,
		DoctorID:         input.DoctorID,
		ListOfDiseasesID: input.ListOfDiseasesID,
		IsOut:            input.IsOut,
		CreatedAt:        input.CreatedAt,
		IsDeleted:        input.IsDeleted,
	}

	models.DB.Create(&dt)

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// PATCH /hospital/:id
// update hospital
func UpdateHospital(c *gin.Context) {
	// get model if exist
	var dt models.HospitalizedHistory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdateHospitalizedInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&dt).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /hospital/delete/:id
// delete a hospital, set IsDeleted as true
func DeleteHospital(c *gin.Context) {
	var dt models.HospitalizedHistory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(UpdateHospitalizedInput{IsDeleted: true})

	c.JSON(http.StatusOK, gin.H{"message": "Hospilatized History deleted successfully"})
}

// RECOVERY

// GET /hospital/deleted
// get deleted hospital
func FindDeletedHospitals(c *gin.Context) {
	var dt []models.HospitalizedHistory
	if err := models.DB.Preload("Patient").Preload("Doctor").Preload("ListOfDiseases").
		Where("is_deleted = ?", true).
		Find(&dt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve hospitals"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /hospital/deleted/:id
// recover deleted hospital
func RecoverDeletedHospital(c *gin.Context) {
	var dt models.HospitalizedHistory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_deleted": false})

	c.JSON(http.StatusOK, gin.H{"data": dt})
}
