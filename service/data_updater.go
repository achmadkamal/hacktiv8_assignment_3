package service

import (
	"hacktiv8_assignment_3/helper"
	"hacktiv8_assignment_3/repository"
	"log"
	"math/rand"
	"time"
)

type DataUpdater struct {
	Repo repository.DataRepository
}

func NewDataUpdater(repo repository.DataRepository) *DataUpdater {
	return &DataUpdater{Repo: repo}
}

func (u *DataUpdater) UpdateData() {
	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		err := u.Repo.UpdateData(water, wind)
		if err != nil {
			log.Println("Error updating data:", err)
			continue
		}

		waterStatus := helper.WaterStatus(water)
		windStatus := helper.WindStatus(wind)

		log.Printf("Water: %d meter, Status: %s", water, waterStatus)
		log.Printf("Wind: %d meter per second, Status: %s", wind, windStatus)

		time.Sleep(15 * time.Second)
	}
}
