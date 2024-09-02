package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	BotToken    string `json:"botToken"`
	PrePrompt   string `json:"prePrompt"`
	Model       string `json:"model"`
	ServerURL   string `json:"serverUrl"`
	EnableLog bool   `json:"enableLog"`
	StartMessage   string `json:"startMessage"`
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Config open error: %v", err)
	}
	defer file.Close()

	config := &Config{}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Config read error: %v", err)
	}

	err = json.Unmarshal(bytes, config)
	if err != nil {
		return nil, fmt.Errorf("Config parse error: %v", err)
	}

	return config, nil
}