package models

import (
	"time"
)

type Patient struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birthdate"`
	Address   string    `json:"address"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type TreatmentHistory struct {
	ID                      string `json:"id" gorm:"primary_key"`
	DiseaseIdentificationAs string `json:"disease_identification_as"`

	// foreign
	PatientID        string         `json:"patient_id"`
	Patient          Patient        `json:"patient" gorm:"foreignKey:PatientID"`
	DoctorID         string         `json:"doctor_id"`
	Doctor           Doctor         `json:"doctor" gorm:"foreignKey:DoctorID"`
	ListOfDiseasesID string         `json:"list_of_diseases_id"`
	ListOfDiseases   ListOfDiseases `json:"list_of_diseases" gorm:"foreignKey:ListOfDiseasesID"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type DrugConsumptionHistory struct {
	ID        string `json:"id" gorm:"primary_key"`
	TotalDose string `json:"total_dose"`
	DoseADay  int    `json:"dose_a_day"`

	// foreign
	PatientID        string         `json:"patient_id"`
	Patient          Patient        `json:"patient" gorm:"foreignKey:PatientID"`
	DoctorID         string         `json:"doctor_id"`
	Doctor           Doctor         `json:"doctor" gorm:"foreignKey:DoctorID"`
	ListOfDiseasesID string         `json:"list_of_diseases_id"`
	ListOfDiseases   ListOfDiseases `json:"list_of_diseases" gorm:"foreignKey:ListOfDiseasesID"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type HospitalizedHistory struct {
	ID string `json:"id" gorm:"primary_key"`

	// foreign
	PatientID        string         `json:"patient_id"`
	Patient          Patient        `json:"patient" gorm:"foreignKey:PatientID"`
	DoctorID         string         `json:"doctor_id"`
	Doctor           Doctor         `json:"doctor" gorm:"foreignKey:DoctorID"`
	ListOfDiseasesID string         `json:"list_of_diseases_id"`
	ListOfDiseases   ListOfDiseases `json:"list_of_diseases" gorm:"foreignKey:ListOfDiseasesID"`

	// status
	IsOut     bool      `json:"is_out"`
	CreatedAt time.Time `json:"created_at"`
	OutAt     time.Time `json:"out_at"`
	IsDeleted bool      `json:"is_deleted"`
}
