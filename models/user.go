package models

import (
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"type:varchar(100);unique_index"`
	Image     string
	Buildings []Building `gorm:"many2many:user_buildings"`
}
