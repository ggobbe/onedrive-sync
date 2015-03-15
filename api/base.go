package api

import (
	"fmt"
	"io/ioutil"
	"log"
)

const rootURL = "https://api.onedrive.com/v1.0"

func doGetRequest(api string) ([]byte, error) {
	URI := fmt.Sprintf("%s%s", rootURL, api)
	log.Printf("GET: %s\n", URI)
	resp, err := client.Get(URI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
