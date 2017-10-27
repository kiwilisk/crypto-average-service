package main

import (
	"os"
	"encoding/json"
)

type Configuration struct {
	AssetEndpointURL string `json:"assetEndpointURL"`
}

func LoadConfiguration() (Configuration, error) {
	file, _ := os.Open("./conf/config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	return configuration, err
}
