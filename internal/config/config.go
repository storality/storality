package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"storality.com/storality/internal/helpers/flags"
)

type Config struct {
	Port     int
	Headless bool
	CMD      string
}

var configFile string = "./storality.config.json"

func init() {
	dataDir := "./stor_data"
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err = os.Mkdir(dataDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func New(cmd string) *Config {
	config, err := load()
	if err != nil {
		config.Port = *flags.Port
		config.Headless = *flags.Headless
	}
	config.CMD = cmd

	if flags.IsPassed("port") {
		config.Port = *flags.Port
	}

	if flags.IsPassed("headless") {
		config.Headless = *flags.Headless
	}

	return &config
}

func load() (Config, error) {
	config := Config{}
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		return config, err
	}
	err = loadExistingConfig(configFile, &config)
	if err != nil {
		log.Fatalf("Error loading existing configuration: %s", err)
	}
	return config, nil
}

func loadExistingConfig(configFile string, config *Config) error {
	file, err := os.Open(configFile)
	if err != nil {
		return fmt.Errorf("failed to open config file: %s", err)
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file '%s': %s", configFile, err)
	}

	err = json.Unmarshal(contents, &config)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON from file '%s': %s", configFile, err)
	}

	return nil
}
