package models

import (
	"database/sql"
	"time"
)

type Administrator struct {
	ID        string       `json:"id" gorm:"primary_key"`
	Name      string       `json:"name"`
	Gender    string       `json:"gender"`
	BirthDate sql.NullTime `json:"birthdate"`
	Address   string       `json:"address"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
