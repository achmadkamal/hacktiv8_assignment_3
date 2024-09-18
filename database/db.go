package database

import (
	"fmt"
	"hacktiv8_assignment_3/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	USER     = "postgres"
	PASSWORD = "postgres"
	PORT     = "5432"
	DATABASE = "hacktiv8_golang_assignment_3"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE, PORT)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Success connected to database")
	db.Debug().AutoMigrate(model.DataModel{})
}

func GetDB() *gorm.DB {
	return db
}
