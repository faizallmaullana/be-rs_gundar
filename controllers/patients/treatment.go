package patients

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/rs_gundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateTreatmentInput struct {
	ID                      string `json:"id" gorm:"primary_key"`
	DiseaseIdentificationAs string `json:"disease_identification_as"`

	// foreign
	PatientID        string `json:"patient_id"`
	Patient          string `json:"patient" gorm:"foreignKey:PatientID"`
	DoctorID         string `json:"doctor_id"`
	Doctor           string `json:"doctor" gorm:"foreignKey:DoctorID"`
	ListOfDiseasesID string `json:"list_of_diseases_id"`
	ListOfDiseases   string `json:"list_of_diseases" gorm:"foreignKey:ListOfDiseasesID"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type UpdateTreatmentInput struct {
	ID                      string `json:"id" gorm:"primary_key"`
	DiseaseIdentificationAs string `json:"disease_identification_as"`

	// foreign
	PatientID        string `json:"patient_id"`
	Patient          string `json:"patient" gorm:"foreignKey:PatientID"`
	DoctorID         string `json:"doctor_id"`
	Doctor           string `json:"doctor" gorm:"foreignKey:DoctorID"`
	ListOfDiseasesID string `json:"list_of_diseases_id"`
	ListOfDiseases   string `json:"list_of_diseases" gorm:"foreignKey:ListOfDiseasesID"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

// TREATMENTS HISTORY

// GET /treatments
// get all treatments
func FindTreatments(c *gin.Context) {
	var dt []models.TreatmentHistory
	if err := models.DB.Preload("Patient").Preload("Doctor").Preload("ListOfDiseases").
		Where("is_deleted = ?", false).
		Find(&dt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve treatments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// GET /treatment/:id
// get a treatment by id
func FindTreatment(c *gin.Context) {
	// get model if it exists
	var dt models.TreatmentHistory
	if err := models.DB.Preload("Patient").Preload("Doctor").Preload("ListOfDiseases").
		Where("id = ? AND is_deleted = ?", c.Param("id"), false).
		First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// POST /treatment
// Create new treatment
func CreateTreatment(c *gin.Context) {
	var input CreateTreatmentInput
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

	// Create treatment
	dt := models.TreatmentHistory{
		ID:                      input.ID,
		DiseaseIdentificationAs: input.DiseaseIdentificationAs,
		PatientID:               input.PatientID,
		DoctorID:                input.DoctorID,
		ListOfDiseasesID:        input.ListOfDiseasesID,
		CreatedAt:               input.CreatedAt,
		IsDeleted:               input.IsDeleted,
	}

	models.DB.Create(&dt)

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// PATCH /treatment/:id
// update treatment
func UpdateTreatment(c *gin.Context) {
	// get model if exist
	var dt models.TreatmentHistory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdateTreatmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&dt).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /treatment/delete/:id
// delete a treatment, set IsDeleted as true
func DeleteTreatment(c *gin.Context) {
	var dt models.TreatmentHistory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(UpdateTreatmentInput{IsDeleted: true})

	c.JSON(http.StatusOK, gin.H{"message": "Treatment History deleted successfully"})
}

// RECOVERY

// GET /treatment/deleted
// get deleted treatment
func FindDeletedTreatments(c *gin.Context) {
	var dt []models.TreatmentHistory
	if err := models.DB.Preload("Patient").Preload("Doctor").Preload("ListOfDiseases").
		Where("is_deleted = ?", true).
		Find(&dt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve treatments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /treatment/deleted/:id
// recover deleted treatment
func RecoverDeletedTreatment(c *gin.Context) {
	var dt models.TreatmentHistory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_deleted": false})

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// GROUP SEARCH

// /treatment/doctor/:doctor_id
func FindTreatmentGDoctor(c *gin.Context) {
	var dt []models.TreatmentHistory
	if err := models.DB.Preload("Patient").Preload("Doctor").Preload("ListOfDiseases").
		Where("doctor_id = ? AND is_deleted = ?", c.Param("doctor_id"), false).
		Find(&dt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve treatments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// /treatment/patient/:patient_id
func FindTreatmentGPatient(c *gin.Context) {
	var dt []models.TreatmentHistory
	if err := models.DB.Preload("Patient").Preload("Doctor").Preload("ListOfDiseases").
		Where("patient_id = ? AND is_deleted = ?", c.Param("patient_id"), false).
		Find(&dt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve treatments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// /treatment/doctor/:id_doctor
func FindTreatmentGListdisease(c *gin.Context) {
	var dt []models.TreatmentHistory
	if err := models.DB.Preload("Patient").Preload("Doctor").Preload("ListOfDiseases").
		Where("list_of_diseases_id = ? AND is_deleted = ?", c.Param("disease_id"), false).
		Find(&dt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve treatments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}
