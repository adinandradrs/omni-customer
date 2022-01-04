package configuration

import (
	"gopkg.in/inconshreveable/log15.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigDatabase() *gorm.DB {
	databaseURL := "postgres://omni-customer:omnicustomerpass@localhost:5432/omni-customer"
	database, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log15.Error("error when try to open database", err)
	}
	return database
}
