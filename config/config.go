package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/metalpoch/go-olt-cantv/model"
)

func LoadConfiguration() model.Config {
	executable_dir, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	base_dir := path.Dir(executable_dir)
	filename := filepath.Join(base_dir, "config.json")

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
