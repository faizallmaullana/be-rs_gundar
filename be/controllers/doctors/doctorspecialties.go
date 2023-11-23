package doctors

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/rs_gundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Docspec represents Doctor Specialties

type CreateDocspecInput struct {
	ID          string `json:"id" gorm:"primary_key"`
	Specialties string `json:"specialties"`
	Description string `json:"description"`

	// status
	IsNew     bool      `json:"is_new"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type UpdateDocspecInput struct {
	ID          string `json:"id" gorm:"primary_key"`
	Specialties string `json:"specialties"`
	Description string `json:"description"`

	// status
	IsNew     bool      `json:"is_new"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

// DOCTOR SPESCIALTIES

// GET /docspecs
// get all Doctor Specialities
func FindDocspecs(c *gin.Context) {
	var dt []models.DoctorSpecialties
	models.DB.Where("is_deleted = ?", false).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// GET /docspec/:id
// get a Doctor Specialities by id
func FindDocspec(c *gin.Context) {
	// get model if its exists
	var dt models.DoctorSpecialties
	if err := models.DB.Where("id = ? AND is_deleted = ?", c.Param("id"), false).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})

}

// POST /docspec
// Create new Doctor Specialties
func CreateDocspec(c *gin.Context) {
	var input CreateDocspecInput
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

	// Auto-set IsNew as true
	input.IsNew = true

	// Create Doctor Specialties
	dt := models.DoctorSpecialties{
		ID:          input.ID,
		Specialties: input.Specialties,
		Description: input.Description,
		IsNew:       input.IsNew,
		CreatedAt:   input.CreatedAt,
		IsDeleted:   input.IsDeleted,
	}

	models.DB.Create(&dt)

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// PATCH /docspec/:id
// update Doctor Specialties
func UpdateDocspec(c *gin.Context) {
	// get model if exist
	var dt models.DoctorSpecialties
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdateDocspecInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&dt).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /docspec/delete/:id
// delete a Doctor Specialties, set IsDeleted as true
func DeleteDocspec(c *gin.Context) {
	var dt models.DoctorSpecialties
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(UpdateDocspecInput{IsDeleted: true})

	c.JSON(http.StatusOK, gin.H{"message": "Doctor Specialties deleted successfully"})
}

// RECOVERY

// GET /docspecs/deleted
// get deleted Doctor Specialties
func FindDeletedDocspecs(c *gin.Context) {
	var dt []models.DoctorSpecialties
	models.DB.Where("is_deleted = ?", true).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /docspec/deleted/:id
// recover deleted Doctor Specialties
func RecoverDeletedDocspec(c *gin.Context) {
	var dt models.DoctorSpecialties
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_deleted": false})

	c.JSON(http.StatusOK, gin.H{"data": dt})
}
