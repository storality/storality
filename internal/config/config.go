package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
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

func Write(port int, headless bool, cmd string) *Config {
	config := &Config{
		Port:     port,
		Headless: headless,
		CMD:      cmd,
	}

	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		saveDefaultConfig(configFile, config)
	} else {
		err = loadExistingConfig(configFile, config)
		if err != nil {
			log.Fatalf("Error loading existing configuration: %s", err)
		}
	}
	return config
}

func Read() *Config {
	config := &Config{}

	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		log.Fatal(err)
	}

	err = loadExistingConfig(configFile, config)
	if err != nil {
		log.Fatalf("Error loading existing configuration: %s", err)
	}

	return config
}

func saveDefaultConfig(configFile string, config *Config) {
	file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	defer file.Close()

	lowercaseConfig := convertToLowerCaseProperties(config)
	configJSON, err := json.MarshalIndent(lowercaseConfig, "", " ")
	if err != nil {
		log.Fatalf("Error marshaling default configuration to JSON: %s", err)
	}

	_, err = file.Write(configJSON)
	if err != nil {
		log.Fatalf("Error writing default configuration to file: %s", err)
	}
}

func convertToLowerCaseProperties(config *Config) interface{} {
	value := reflect.ValueOf(config).Elem()
	typeOfValue := value.Type()

	lowercaseConfig := make(map[string]interface{})
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := strings.ToLower(typeOfValue.Field(i).Name)

		lowercaseConfig[fieldName] = field.Interface()
	}

	return lowercaseConfig
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
