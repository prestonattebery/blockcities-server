package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Name      string     `gorm:"unique;not null"`
	Username  string     `gorm:"unique;not null"`
	Email     string     `gorm:"type:varchar(100);unique_index"`
	Buildings []Building `gorm:"many2many:Build;association_foreignkey:address;foreignkey:idPerson"`
}
