package entity

import(
	"gorm.io/gorm"
)

type Tests struct {
	gorm.Model
	Name string `valid:"required~testNO, email"`
}