package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/metalpoch/go-olt-cantv/model"
)

func LoadConfiguration() model.Config {
	var config model.Config

	f, err := os.ReadFile("config.json")
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(f), &config)

	return config
}
