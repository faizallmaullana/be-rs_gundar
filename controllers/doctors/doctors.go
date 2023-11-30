package doctors

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/rs_gundar/models"
	"github.com/gin-gonic/gin"
)

type CreateDoctorInput struct {
	ID        string `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Gender    bool   `json:"gender"`
	BirthDate string `json:"birthdate"`
	Address   string `json:"address"`

	// foreign
	DoctorSpecialitiesID string `json:"doctor_specialities_id"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type UpdateDoctorInput struct {
	ID        string `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Gender    bool   `json:"gender"`
	BirthDate string `json:"birthdate"`
	Address   string `json:"address"`

	// foreign
	DoctorSpecialitiesID string `json:"doctor_specialities_id"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

// DOCTORS

// GET /doctors
// get all doctors
func FindDoctors(c *gin.Context) {
	var dt []models.Doctor
	models.DB.Where("is_deleted = ?", false).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// GET /doctor/:id
// get a doctor by id
func FindDoctor(c *gin.Context) {
	// get model if its exists
	var dt models.Doctor
	if err := models.DB.Where("id = ? AND is_deleted = ?", c.Param("id"), false).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dt})

}

// POST /doctor
// Create new doctor
func CreateDoctor(c *gin.Context) {
	var input CreateDoctorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// id doctor is nomor induk profesi dokter

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

	// Create doctor
	dt := models.Doctor{
		ID:                   input.ID,
		Name:                 input.Name,
		Gender:               input.Gender,
		BirthDate:            parsedTime,
		Address:              input.Address,
		DoctorSpecialitiesID: input.DoctorSpecialitiesID,
		CreatedAt:            input.CreatedAt,
		IsDeleted:            input.IsDeleted,
	}

	models.DB.Create(&dt)

	c.JSON(http.StatusCreated, gin.H{"data": dt})
}

// PATCH /doctor/:id
// update doctor
func UpdateDoctor(c *gin.Context) {
	// get model if exist
	var dt models.Doctor
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdateDoctorInput
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

	data := models.Doctor{
		Name:                 input.Name,
		BirthDate:            parsedTime,
		Gender:               input.Gender,
		Address:              input.Address,
		DoctorSpecialitiesID: input.DoctorSpecialitiesID,
	}

	models.DB.Model(&dt).Updates(data)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /doctor/delete/:id
// delete a doctor, set IsDeleted as true
func DeleteDoctor(c *gin.Context) {
	var dt models.Doctor
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(UpdateDoctorInput{IsDeleted: true})

	c.JSON(http.StatusOK, gin.H{"message": "Doctor deleted successfully"})
}

// RECOVERY

// GET /doctor/deleted
// get deleted doctor
func FindDeletedDoctor(c *gin.Context) {
	var dt []models.Doctor
	models.DB.Where("is_deleted = ?", true).Find(&dt)

	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// DELETE /doctor/deleted/:id
// recover deleted doctor
func RecoverDeletedDoctor(c *gin.Context) {
	var dt models.Doctor
	if err := models.DB.Where("id = ?", c.Param("id")).First(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&dt).Updates(map[string]interface{}{"is_deleted": false})

	c.JSON(http.StatusOK, gin.H{"data": dt})
}
