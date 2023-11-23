package models

import (
	"time"
)

type Administrator struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birthdate"`
	Address   string    `json:"address"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
