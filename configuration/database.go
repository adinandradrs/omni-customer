package configuration

import (
	"omni-customer/model/entity"

	"gopkg.in/inconshreveable/log15.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigDatabase(databaseConfig string) *gorm.DB {
	database, err := gorm.Open(postgres.Open(databaseConfig), &gorm.Config{})
	if err != nil {
		log15.Error("error when try to open database", err)
	}
	database.AutoMigrate(&entity.Customer{})
	return database
}
