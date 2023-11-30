package models

import (
	"time"
)

type LoginDoctor struct {
	ID       string `json:"id" gorm:"primary_key"`
	Password string `json:"password"`
}

type Doctor struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Gender    bool      `json:"gender"`
	BirthDate time.Time `json:"birthdate"`
	Address   string    `json:"address"`

	// foreignkey
	DoctorSpecialitiesID string            `json:"doctor_specialities_id"`
	DoctorSpecialties    DoctorSpecialties `json:"doctor_specialities" gorm:"foreignKey:DoctorSpecialtiesID"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type DoctorSpecialties struct {
	ID          string `json:"id" gorm:"primary_key"`
	Specialties string `json:"specialties"`
	Description string `json:"description"`

	// status
	IsNew     bool      `json:"is_new"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
