package database

import "api/models"

func Init() {

	item := &models.Items{name: "test", serial: 512135435}

	Database.Db.Create(&item)
}
