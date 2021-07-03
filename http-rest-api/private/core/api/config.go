package api

import (
	"github.com/Survereignty/spet-rest-api/private/core/api/db"
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config ...
type Config struct {
	Address string `json:"Address"`
	DataBase *db.Config
}

const path = "./config/config.json"

// NewConfig ...
func NewConfig(с *Config) (*Config, error) {

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {

			// Нет, значит создаем
			reload(с)

		} else {
			return nil, err
		}
	}

	dataJSON, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data := &Config{}

	err = json.Unmarshal([]byte(dataJSON), data)
	if err != nil {
		reload(с)
		dataJSON, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(dataJSON), data)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func reload(c *Config) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil
	}

	defer f.Close()

	dataJSON, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, dataJSON, 0)
	if err != nil {
		return err
	}
	return nil
}
