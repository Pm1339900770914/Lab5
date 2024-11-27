package entity

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Username string `gorm:"uniqueIndex" valid:"required~Username is required"`
	Password string 
	Email  string `gorm:"uniqueIndex" valid:"required~Email is required, email~Email is invalid"`

}
