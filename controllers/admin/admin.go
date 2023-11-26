package admin

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/rs_gundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateAdminInput struct {
	ID        string `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthdate"`
	Address   string `json:"address"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type UpdateAdminInput struct {
	ID        string `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthdate"`
	Address   string `json:"address"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

// ADMINISTRATOR

// GET /admins
// get all admins
func FindAdmins(c *gin.Context) {
	var dt []models.Administrator
	models.DB.Where("is_deleted = ?", false).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// GET /admin/:id
// get a admin by id
func FindAdmin(c *gin.Context) {
	// get model if its exists
	var dt models.Administrator
	if err := models.DB.Where("id = ? AND is_deleted = ?", c.Param("id"), false).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})

}

// POST /admin
// Create new admin
func CreateAdmin(c *gin.Context) {
	var input CreateAdminInput
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

	// Create Admin
	dt := models.Administrator{
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

// PATCH /admin/:id
// update admin
func UpdateAdmin(c *gin.Context) {
	// get model if exist
	var dt models.Administrator
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdateAdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&dt).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /admin/delete/:id
// delete a admin, set IsDeleted as true
func DeleteAdmin(c *gin.Context) {
	var dt models.Administrator
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(UpdateAdminInput{IsDeleted: true})

	c.JSON(http.StatusOK, gin.H{"message": "Admin deleted successfully"})
}

// RECOVERY

// GET /admins/deleted
// get deleted admin
func FindDeletedAdmins(c *gin.Context) {
	var dt []models.Administrator
	models.DB.Where("is_deleted = ?", true).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /admin/deleted/:id
// recover deleted admin
func RecoverDeletedAdmin(c *gin.Context) {
	var dt models.Administrator
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_deleted": false})

	c.JSON(http.StatusOK, gin.H{"data": dt})
}
