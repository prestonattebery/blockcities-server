package models

import "github.com/jinzhu/gorm"

type Building struct {
	gorm.Model

	Title   string
	Address uint
}
