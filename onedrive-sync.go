package main

import (
	"fmt"
	"log"

	"github.com/ggobbe/onedrive-sync/api"
)

func main() {
	log.Println("Starting")

	drive, err := api.GetDrive()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Drive ID: %s\n", drive.ID)

	root, err := api.GetRoot()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Root Size: %d\n", root.Size)

	children, err := api.GetChildren()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Children: %d\n", len(children.Value))

	log.Println("Finished")
}
