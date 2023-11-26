package listdiseases

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/rs_gundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateDrugsInput struct {
	ID            string `json:"id" gorm:"primary_key"`
	Drug          string `json:"drug"`
	Description   string `json:"description"`
	Clasification string `json:"clasification"` // list are set in the controller (obat bebas, bebas terbatas, keras, narkotika)

	// stauts
	IsNew     bool      `json:"is_new"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type UpdateDrugsInput struct {
	ID            string `json:"id" gorm:"primary_key"`
	Drug          string `json:"drug"`
	Description   string `json:"description"`
	Clasification string `json:"clasification"` // list are set in the controller (obat bebas, bebas terbatas, keras, narkotika)

	// stauts
	IsNew     bool      `json:"is_new"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

// List Of DRUGS

// GET /drugs
// get all drugs information
func FindDrugs(c *gin.Context) {
	var dt []models.ListOfDrugs
	models.DB.Where("is_deleted = ?", false).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// GET /drug/:id
// get drug by id
func FindDrug(c *gin.Context) {
	var dt models.ListOfDrugs
	if err := models.DB.Where("id = ? AND is_deleted = ?", c.Param("id"), false).First(&dt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// POST /drug
// create new drug list
func CreateDrug(c *gin.Context) {
	var input CreateDrugsInput
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

	// create new drug
	dt := models.ListOfDrugs{
		ID:            input.ID,
		Drug:          input.Drug,
		Description:   input.Description,
		Clasification: input.Clasification,
		IsNew:         input.IsNew,
		CreatedAt:     input.CreatedAt,
		IsDeleted:     input.IsDeleted,
	}

	models.DB.Create(&dt)

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// PATCH /drug/:id
// update drug
func UpdateDrug(c *gin.Context) {
	var dt models.ListOfDrugs
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdateDrugsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&dt).Updates(input)

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_new": false})

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// DELETE /drug/delete/:id
// delete a patient, set IsDeleted as true
func DeleteDrug(c *gin.Context) {
	var dt models.ListOfDrugs
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(UpdateDrugsInput{IsDeleted: true})

	c.JSON(http.StatusOK, gin.H{"message": "Drug deleted successfully"})
}

// RECOVERY

// GET /drugs/deleted
// get deleted drug
func FindDeletedDrugs(c *gin.Context) {
	var dt []models.ListOfDrugs
	models.DB.Where("is_deleted = ?", true).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /drug/deleted/:id
// recover deleted patient
func RecoverDeletedDrug(c *gin.Context) {
	var dt models.ListOfDrugs
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_deleted": false})

	c.JSON(http.StatusOK, gin.H{"data": dt})
}
