package main

import (
	"hacktiv8_assignment_3/database"
	"hacktiv8_assignment_3/repository"
	"hacktiv8_assignment_3/service"
	"log"
)

func main() {
	database.StartDB()

	db := database.GetDB()
	dataRepo := repository.NewDataRepository(db)
	updater := service.NewDataUpdater(dataRepo)

	go updater.UpdateData()

	log.Println("Service is running...")
	select {}
}
