package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/oauth2"
)

const configName = ".onedrive-sync"
const tokenName = "token"

func init() {
	configPath, err := getConfigPath()
	if err != nil {
		log.Fatal(err)
	}
	err = createIfNotExists(configPath)
	if err != nil {
		log.Fatal(err)
	}
}

// SaveToken saves the token on the disk
func SaveToken(token *oauth2.Token) error {
	content, err := json.Marshal(token)
	if err != nil {
		return err
	}
	tokenFile, err := getTokenPath()
	if err != nil {
		return err
	}
	log.Printf("Saving token in %s\n", tokenFile)
	return ioutil.WriteFile(tokenFile, content, 0644)
}

// ReadToken returns the token saved on the disk
func ReadToken() (*oauth2.Token, error) {
	tokenFile, err := getTokenPath()
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return nil, err
	}
	var token *oauth2.Token
	err = json.Unmarshal(content, &token)
	return token, err
}

func getConfigPath() (string, error) {
	userHomeDir, err := getUserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s%s", userHomeDir, string(filepath.Separator), configName), nil
}

func getTokenPath() (string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s%s", configPath, string(filepath.Separator), tokenName), nil
}

func getUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

func createIfNotExists(dir string) error {
	exists, err := exists(dir)
	if err != nil || exists {
		return err
	}
	log.Printf("Creating folder %s\n", dir)
	err = os.Mkdir(dir, 0777)
	return err
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
