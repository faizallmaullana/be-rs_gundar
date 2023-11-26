package listdiseases

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/rs_gundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateChamberInput struct {
	ID          string `json:"id" gorm:"primary_key"`
	ChamberName string `json:"chamber_name"`
	Capacity    int    `json:"capacity"`
	// for see the filled capacity, substract hospitalization chamber that status of isOut 1 on the patient
	// capacity - isOut

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type UpdateChamberInput struct {
	ID          string `json:"id" gorm:"primary_key"`
	ChamberName string `json:"chamber_name"`
	Capacity    int    `json:"capacity"`
	// for see the filled capacity, substract hospitalization chamber that status of isOut 1 on the patient
	// capacity - isOut

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

// List Of Chamber

// GET /chambers
// get all chambers information
func FindChambers(c *gin.Context) {
	var dt []models.ListOfHospitalChambers
	models.DB.Where("is_deleted = ?", false).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// GET /chamber/:id
// get chamber by id
func FindChamber(c *gin.Context) {
	var dt models.ListOfHospitalChambers
	if err := models.DB.Where("id = ? AND is_deleted = ?", c.Param("id"), false).First(&dt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// POST /chamber
// create new chamber list
func CreateChamber(c *gin.Context) {
	var input CreateChamberInput
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

	// create new chamber
	dt := models.ListOfHospitalChambers{
		ID:          input.ID,
		ChamberName: input.ChamberName,
		Capacity:    input.Capacity,
		CreatedAt:   input.CreatedAt,
		IsDeleted:   input.IsDeleted,
	}

	models.DB.Create(&dt)

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// PATCH /chamber/:id
// update chamber
func UpdateChamber(c *gin.Context) {
	var dt models.ListOfHospitalChambers
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdateChamberInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&dt).Updates(input)

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_new": false})

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// DELETE /chamber/delete/:id
// delete a patient, set IsDeleted as true
func DeleteChamber(c *gin.Context) {
	var dt models.ListOfHospitalChambers
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(UpdateChamberInput{IsDeleted: true})

	c.JSON(http.StatusOK, gin.H{"message": "Chamber deleted successfully"})
}

// RECOVERY

// GET /chambers/deleted
// get deleted chamber
func FindDeletedChambers(c *gin.Context) {
	var dt []models.ListOfHospitalChambers
	models.DB.Where("is_deleted = ?", true).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /chamber/deleted/:id
// recover deleted patient
func RecoverDeletedChamber(c *gin.Context) {
	var dt models.ListOfHospitalChambers
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_deleted": false})

	c.JSON(http.StatusOK, gin.H{"data": dt})
}
