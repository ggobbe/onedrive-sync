package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// Drive type
type Drive struct {
	DataContext string `json:"@odata.context"`
	ID          string `json:"id"`
	DriveType   string `json:"driveType"`
	Owner       Owner  `json:"owner"`
	Quota       Quota  `json:"quota"`
}

// Owner type
type Owner struct {
	User User `json:"user"`
}

// User type
type User struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
}

// Quota type
type Quota struct {
	Deleted   int    `json:"deleted"`
	Remaining int    `json:"remaining"`
	State     string `json:"state"`
	Total     int    `json:"total"`
	Used      int    `json:"used"`
}

// Error type
type Error struct {
	Code       string     `json:"code"`
	Message    string     `json:"message"`
	InnerError InnerError `json:"innererror"`
}

// InnerError type
type InnerError struct {
	Code string `json:"code"`
}

// GetDriveInfos return the drive information
func GetDriveInfos() (*Drive, error) {
	driveURI := fmt.Sprintf("%s/drive/", apiURL)
	log.Printf("Request: %s\n", driveURI)
	resp, err := client.Get(driveURI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var data Drive
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
