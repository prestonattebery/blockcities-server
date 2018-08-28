package models

import "github.com/jinzhu/gorm"

// Building Struct
type Building struct {
	gorm.Model

	Name      string `gorm:"not null"`
	Summary   string `gorm:"index"`
	Built     string
	Height    string
	Architect string
	ImageURL  string
}
