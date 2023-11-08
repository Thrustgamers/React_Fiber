package models

import (
	"gorm.io/gorm"
)

type Users struct {
	rank       Ranks `gorm:"embedded"`
	name       string
	employeeId int16
	password   string
	gorm.Model
}
