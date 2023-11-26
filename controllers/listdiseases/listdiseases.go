package listdiseases

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/rs_gundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateDiseaseInput struct {
	ID          string `json:"id" gorm:"primary_key"`
	Disease     string `json:"disease"` // buat satu disease dengan nama kondisi baik
	Description string `json:"description"`
	Infectious  string `json:"infectious"` // apakah menular

	// status
	IsNew     bool      `json:"is_new"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type UpdateDiseaseInput struct {
	ID          string `json:"id" gorm:"primary_key"`
	Disease     string `json:"disease"` // buat satu disease dengan nama kondisi baik
	Description string `json:"description"`
	Infectious  string `json:"infectious"` // apakah menular

	// status
	IsNew     bool      `json:"is_new"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

// List Of Disease

// GET /diseases
// get all disease information
func FindDiseases(c *gin.Context) {
	var dt []models.ListOfDiseases
	models.DB.Where("is_deleted = ?", false).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// GET /disease/:id
// get disease by id
func FindDisease(c *gin.Context) {
	var dt models.ListOfDiseases
	if err := models.DB.Where("id = ? AND is_deleted = ?", c.Param("id"), false).First(&dt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// POST /disease
// create new disease list
func CreateDisease(c *gin.Context) {
	var input CreateDiseaseInput
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

	// auto-set isNew to true
	input.IsNew = true

	// create new disease
	dt := models.ListOfDiseases{
		ID:          input.ID,
		Disease:     input.Disease,
		Description: input.Description,
		Infectious:  input.Infectious,
		IsNew:       input.IsNew,
		CreatedAt:   input.CreatedAt,
		IsDeleted:   input.IsDeleted,
	}

	models.DB.Create(&dt)

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// PATCH /disease/:id
// update disease
func UpdateDisease(c *gin.Context) {
	var dt models.ListOfDiseases
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdateDiseaseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&dt).Updates(input)

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_new": false})

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// DELETE /disease/delete/:id
// delete a patient, set IsDeleted as true
func DeleteDisease(c *gin.Context) {
	var dt models.ListOfDiseases
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(UpdateDiseaseInput{IsDeleted: true})

	c.JSON(http.StatusOK, gin.H{"message": "Disease deleted successfully"})
}

// RECOVERY

// GET /diseases/deleted
// get deleted disease
func FindDeletedDiseases(c *gin.Context) {
	var dt []models.ListOfDiseases
	models.DB.Where("is_deleted = ?", true).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /disease/deleted/:id
// recover deleted patient
func RecoverDeletedDisease(c *gin.Context) {
	var dt models.ListOfDiseases
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_deleted": false})

	c.JSON(http.StatusOK, gin.H{"data": dt})
}
