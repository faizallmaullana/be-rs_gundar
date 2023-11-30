package models

import "time"

type ListOfDiseases struct {
	ID          string `json:"id" gorm:"primary_key"`
	Disease     string `json:"disease"` // buat satu disease dengan nama kondisi baik
	Description string `json:"description"`
	Infectious  bool   `json:"infectious"` // apakah menular

	// status
	IsNew     bool      `json:"is_new"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type ListOfDrugs struct {
	ID            string `json:"id" gorm:"primary_key"`
	Drug          string `json:"drug"`
	Description   string `json:"description"`
	Clasification string `json:"clasification"` // list are set in the controller (obat bebas, bebas terbatas, keras, narkotika)

	// stauts
	IsNew     bool      `json:"is_new"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type ListOfHospitalChambers struct {
	ID          string `json:"id" gorm:"primary_key"`
	ChamberName string `json:"chamber_name"`
	Capacity    int    `json:"capacity"`
	// for see the filled capacity, substract hospitalization chamber that status of isOut 1 on the patient
	// capacity - isOut

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
