package models

import (
	"gorm.io/gorm"
)

type Items struct {
	name   string `json:"name"`
	serial int64  `json:"serial"`
	gorm.Model
}
