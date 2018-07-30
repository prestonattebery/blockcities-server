package models

import "github.com/jinzhu/gorm"

// Building Struct
type Building struct {
	gorm.Model

	Title   string `gorm:"not null"`
	Address string `gorm:"index"`
}
