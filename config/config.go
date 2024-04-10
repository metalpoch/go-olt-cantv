package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/metalpoch/go-olt-cantv/model"
)

func LoadConfiguration() model.Config {
	directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filename := filepath.Join(directory, "config.json")

	var config model.Config
	configFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
