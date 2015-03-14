package main

import (
	"log"

	"github.com/ggobbe/onedrive-sync/api"
)

func main() {
	log.Println("Starting onesync")
	drive, err := api.GetDriveInfos()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Drive ID: %s\n", drive.ID)
	log.Println("Finished")
}
