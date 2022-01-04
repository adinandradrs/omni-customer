package controller

import "gorm.io/gorm"

type DatabaseHandler struct {
	DB *gorm.DB
}

func RestRepository(database *gorm.DB) DatabaseHandler {
	return DatabaseHandler{database}
}
