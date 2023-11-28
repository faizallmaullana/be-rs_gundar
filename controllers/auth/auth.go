package auth

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/rs_gundar/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthInput struct {
	ID        string    `json:"id" grom:"primary_key"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

// POST /auth/admin/register
// Create new admin
func RegisterToAdmin(c *gin.Context) {
	var input AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the username already exists
	var existingAdmin models.LoginAdministrator
	if err := models.DB.Where("username = ?", input.Username).First(&existingAdmin).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// cek token
	if input.Token != "tokenadmin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token salah"})
		return
	}

	// Auto-generate ID using UUID
	input.ID = uuid.New().String()

	// Auto-generate CreatedAt to UTC+7
	input.CreatedAt = time.Now().UTC().Add(7 * time.Hour)

	// Auto-set IsDeleted as 0
	input.IsDeleted = false

	// Create Admin
	dt := models.LoginAdministrator{
		ID:       input.ID,
		Username: input.Username,
		Password: input.Password,
	}

	dt2 := models.Administrator{
		ID:        input.ID,
		CreatedAt: input.CreatedAt,
		IsDeleted: input.IsDeleted,
	}

	models.DB.Create(&dt)
	models.DB.Create(&dt2)

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// POST /auth/admin/login
// Admin login
func LoginToAdmin(c *gin.Context) {
	var input AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the admin by username
	var admin models.LoginAdministrator
	if err := models.DB.Where("username = ?", input.Username).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if the provided password is correct
	if admin.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin login successful", "data": admin})
}

// POST /auth/admin/register
func RegisterToDoctor(c *gin.Context) {
	var input AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek token
	if input.Token != "tokendoctor" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token salah"})
		return
	}

	// Auto-generate CreatedAt to UTC+7
	input.CreatedAt = time.Now().UTC().Add(7 * time.Hour)

	// Auto-set IsDeleted as 0
	input.IsDeleted = false

	// Create Admin
	dt := models.LoginDoctor{
		ID:       input.Username,
		Password: input.Password,
	}

	dt2 := models.Doctor{
		ID:        input.Username,
		CreatedAt: input.CreatedAt,
		IsDeleted: input.IsDeleted,
	}

	models.DB.Create(&dt)
	models.DB.Create(&dt2)

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// POST /auth/doctor/login
// Doctor login
func LoginToDoctor(c *gin.Context) {
	var input AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the doctor by username
	var doctor models.LoginDoctor
	if err := models.DB.Where("id = ?", input.Username).First(&doctor).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if the provided password is correct
	if doctor.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor login successful", "data": doctor})
}
