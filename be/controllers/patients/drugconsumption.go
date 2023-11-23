package patients

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/rs_gundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//  Drugcons represent Drug Consumsion History

type CreateDrugconsHistory struct {
	ID        string `json:"id" gorm:"primary_key"`
	TotalDose string `json:"total_dose"`
	DoseADay  int    `json:"dose_a_day"`

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
type UpdateDrugconsInput struct {
	ID        string `json:"id" gorm:"primary_key"`
	TotalDose string `json:"total_dose"`
	DoseADay  int    `json:"dose_a_day"`

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

// DRUG CONSUMPTION

// GET /drugconss
// get all drugconss
func FindDrugconss(c *gin.Context) {
	var dt []models.DrugConsumptionHistory
	if err := models.DB.Preload("Patient").Preload("Doctor").Preload("ListOfDiseases").
		Where("is_deleted = ?", false).
		Find(&dt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve drugconss"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// GET /drugcons/:id
// get a drugcons by id
func FindDrugcons(c *gin.Context) {
	// get model if it exists
	var dt models.DrugConsumptionHistory
	if err := models.DB.Preload("Patient").Preload("Doctor").Preload("ListOfDiseases").
		Where("id = ? AND is_deleted = ?", c.Param("id"), false).
		First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// POST /drugcons
// Create new drugcons
func CreateDrugcons(c *gin.Context) {
	var input CreateDrugconsHistory
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

	// Create drugcons
	dt := models.DrugConsumptionHistory{
		ID:        input.ID,
		TotalDose: input.TotalDose,
		PatientID: input.PatientID,
		DoctorID:  input.DoctorID,
		CreatedAt: input.CreatedAt,
		IsDeleted: input.IsDeleted,
	}

	models.DB.Create(&dt)

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// PATCH /drugcons/:id
// update drugcons
func UpdateDrugcons(c *gin.Context) {
	// get model if exist
	var dt models.DrugConsumptionHistory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdateDrugconsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&dt).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /drugcons/delete/:id
// delete a drugcons, set IsDeleted as true
func DeleteDrugcons(c *gin.Context) {
	var dt models.DrugConsumptionHistory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(UpdateDrugconsInput{IsDeleted: true})

	c.JSON(http.StatusOK, gin.H{"message": "Drugs Consumtion History deleted successfully"})
}

// RECOVERY

// GET /drugcons/deleted
// get deleted drugcons
func FindDeletedDrugconss(c *gin.Context) {
	var dt []models.DrugConsumptionHistory
	if err := models.DB.Preload("Patient").Preload("Doctor").Preload("ListOfDiseases").
		Where("is_deleted = ?", true).
		Find(&dt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve drugconss"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /drugcons/deleted/:id
// recover deleted drugcons
func RecoverDeletedDrugcons(c *gin.Context) {
	var dt models.DrugConsumptionHistory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_deleted": false})

	c.JSON(http.StatusOK, gin.H{"data": dt})
}
