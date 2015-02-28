package api

import (
	"fmt"
	"log"
	"net/http"
)

const apiURL string = "https://api.onedrive.com/v1.0"

var client *http.Client

func init() {
	c, err := getClient()
	if err != nil {
		log.Fatal(err)
	}
	client = c
}

// GetDriveInfos return the content length of the /drive/ api entry point
func GetDriveInfos() string {
	driveURI := fmt.Sprintf("%s/drive/", apiURL)
	log.Printf("Request: %s\n", driveURI)
	resp, err := client.Get(driveURI)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Status: %s\nContentLength: %d", resp.Status, resp.ContentLength)
}
