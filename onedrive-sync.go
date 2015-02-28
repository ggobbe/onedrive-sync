package main

import (
	"fmt"
	"log"

	"github.com/ggobbe/onedrive-sync/api"
)

func main() {
	log.Println("Starting onesync")

	fmt.Println(api.GetDriveInfos())

	log.Println("Finished")
}
