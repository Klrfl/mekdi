package models

import "github.com/google/uuid"

type Menu struct {
	ID          uuid.UUID
	Name        string
	ServingSize string
	Ingredients string
	Tag         string
	Allergy     string
	Energy      float64
	Protein     float64
	TotalFat    float64
	SatFat      float64
	TransFat    float64
	Chol        float64
	Carbs       float64
	TotalSugar  float64
	AddedSugar  float64
	Sodium      float64
	Description string
}
