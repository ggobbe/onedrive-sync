package api

import (
	"fmt"
	"log"
	"net/http"
)

const apiURL string = "https://api.onedrive.com/v1.0/"

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
	resp, err := client.Get(fmt.Sprintf("%s/drive/", apiURL))
	if err != nil {
		log.Fatalln(err)
	}
	return fmt.Sprintf("Content Length: %d", resp.ContentLength)
}
