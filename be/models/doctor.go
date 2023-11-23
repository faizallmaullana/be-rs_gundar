package models

import (
	"time"
)

type Doctor struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birthdate"`
	Address   string    `json:"address"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type ListOfSpecialties struct {
	ID          string `json:"id" gorm:"primary_key"`
	Specialties string `json:"specialties"`
	Description string `json:"description"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
