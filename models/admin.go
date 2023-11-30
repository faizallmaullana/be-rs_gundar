package models

import (
	"time"
)

type LoginAdministrator struct {
	Username string `json:"username"`
	Password string `json:"password"`

	// foreignkey
	AdministratorID string        `json:"administrator_id"`
	Administrator   Administrator `json:"administrator" gorm:"foreignKey:AdministratorID"`
}

type Administrator struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Gender    bool      `json:"gender"`
	BirthDate time.Time `json:"birthdate"`
	Address   string    `json:"address"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
