package models

import (
	"gorm.io/gorm"
)

type Ranks struct {
	name string
	gorm.Model
}
