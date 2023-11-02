package models

type Users struct {
	ID         uint
	rank       Ranks `gorm:"embedded"`
	name       string
	employeeId int16
}
