package core

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ServerAddress struct {
	HostName string `yaml:"host"`
	Port     string `yaml:"port"`
}

type TLS struct {
	CrtPath string `yaml:"crt"`
	Key     string `yaml:"key"`
}

type ConnectionCredentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
}

type Config struct {
	ServerAddress ServerAddress          `yaml:"server"`
	Tls           *TLS                   `yaml:"tls,omitempty"`
	AccessLogging *ConnectionCredentials `yaml:"logging_server,omitempty"`
}

// fileExists checks if a file exists and is not a directory before we try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// LoadConfigFromFile loads the Config from the given path
func LoadConfigFromFile(path string) (Config, error) {
	if fileExists(path) {
		content, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error while reading file %s: %v", path, err)
			return Config{}, err
		}
		var config Config
		err = yaml.Unmarshal(content, &config)
		if err != nil {
			log.Printf("Error parsing YAML file %s: %v", path, err)
		}
		return config, nil
	}
	return Config{}, fmt.Errorf("%s does not exists", path)
}
