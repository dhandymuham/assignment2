package config

import (
	"fmt"
	"test/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "koinworks"
	dbname   = "orders_by"
	db       *gorm.DB
	err      error
)

func InitDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(models.Orders{}, models.Items{})

	return db

}
