package config

import (
	"github.com/m1keru/pushbot/internal/alertmanager"
	log "github.com/sirupsen/logrus"
	"os"

	"gopkg.in/yaml.v2"
)

// Config -- Global Config
type Config struct {
	Daemon       Daemon                    `yaml:"daemon"`
	Alertmanager alertmanager.Alertmanager `yaml:"alertmanager"`
	Databases    []Database                `yaml:"databases"`
}

// Database -- databases
type Database struct {
	Type        string `yaml:"type"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	FailTimeout int    `yaml:"fail_timeout"`
	Schema      string `yaml:"schema,omitempty"`
}

//Daemon - DaemonConfig
type Daemon struct {
	LogFile string `yaml:"LogFile"`
	Debug   bool   `yaml:"Debug"`
	Port    int    `yaml:"port"`
}

//Setup - Setup
func (cfg *Config) Setup(filename *string) error {
	configFile, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Unable to read config file. Error:\n %v\n", err)
	}
	defer configFile.Close()
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("Unable to Unmarshal Config, Error:\n %v\n", err)
	}
	return nil
}
