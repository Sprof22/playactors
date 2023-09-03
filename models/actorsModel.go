package models

import "gorm.io/gorm"

type Actors struct {
	gorm.Model
	ActorName       string
	ActorRating     int
	ImagePath       string
	AlternativeName string
	ActorID         int
}
