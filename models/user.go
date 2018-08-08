package models

import (
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model

	Name      string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"type:varchar(100);unique_index"`
	Image     string
	Buildings []Building `gorm:"many2many:user_buildings"`
}
