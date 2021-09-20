package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ClientConfig struct {
	Hosts []string
}

var Config ClientConfig

func ReadConfig(configFilePath string) error {
	file, err := os.Open(configFilePath)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Config)
	if err != nil {
		return err
	}

	return nil
}
