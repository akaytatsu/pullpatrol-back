package entity

import "gorm.io/gorm"

type EntityGroup struct {
	gorm.Model
	Name        string
	Description string
}
