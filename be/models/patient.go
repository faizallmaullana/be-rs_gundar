package models

import (
	"database/sql"
	"time"
)

type Patient struct {
	ID        string       `json:"id" gorm:"primary_key"`
	Name      string       `json:"name"`
	Gender    string       `json:"gender"`
	BirthDate sql.NullTime `json:"birthdate"`
	Address   string       `json:"address"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type DiseaseHistory struct {
	ID string `json:"id" gorm:"primary_key"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`

	ListOfDiseasesID string         `json:"list_of_diseases_id"`
	ListOfDiseases   ListOfDiseases `json:"list_of_diseases_list" gorm:"foreignKey:ListOfDiseasesID"`
}

type TreatmentHistory struct {
	ID string `json:"id" gorm:"primary_key"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type DrugConsumptionHistory struct {
	ID string `json:"id" gorm:"primary_key"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type HospitalizedHistory struct {
	ID string `json:"id" gorm:"primary_key"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type ListOfDiseases struct {
	ID string `json:"id" gorm:"primary_key"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type ListofDrugs struct {
	ID string `json:"id" gorm:"primary_key"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type ListOfHospitalChabers struct {
	ID string `json:"id" gorm:"primary_key"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type ListOfTreatments struct {
	ID string `json:"id" gorm:"primary_key"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
