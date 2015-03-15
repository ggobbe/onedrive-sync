package api

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	driveAPI    = "/drive/"
	rootAPI     = "/drive/root"
	childrenAPI = "/drive/root/children"
)

var client *http.Client

func init() {
	c, err := getClient()
	if err != nil {
		log.Fatal(err)
	}
	client = c
}

// GetDrive return the drive information
func GetDrive() (*Drive, error) {
	body, err := doGetRequest(driveAPI)
	if err != nil {
		return nil, err
	}
	var data Drive
	err = json.Unmarshal(body, &data)
	return &data, err
}

// GetRoot gets the root folder for the user's default Drive
func GetRoot() (Item, error) {
	body, err := doGetRequest(rootAPI)
	if err != nil {
		log.Fatal(err)
	}
	var data Item
	err = json.Unmarshal(body, &data)
	return data, err
}

// GetChildren List children under the Drive
func GetChildren() (Children, error) {
	body, err := doGetRequest(childrenAPI)
	if err != nil {
		log.Fatal(err)
	}
	var data Children
	err = json.Unmarshal(body, &data)
	return data, err
}
