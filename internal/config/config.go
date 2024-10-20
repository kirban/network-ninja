package config

import (
	"os"
	"log"

	yaml "gopkg.in/yaml.v2"
	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig `yaml:"app"`
	Resources []ResourcesList `yaml:"resources"`
	Alerts AlertsConfig `yaml:"alerts"`
	Logs LogsConfig `yaml:"logs"`
	Grafana GrafanaConfig `yaml:"grafana"`
}

func (c *Config) Load() *Config {
	envFile, err := godotenv.Read(".env")

	path := envFile["CONFIG_PATH"]

	if err != nil {
		log.Printf("Failed to parse env variables: #%v ", err)
	}

	configFile, err := os.ReadFile(path)

	if err != nil {
		log.Println(path)
		log.Printf("Failed to read config file: #%v ", err)
	}

	errUnmarshal := yaml.Unmarshal([]byte(configFile), c)

	if errUnmarshal != nil {
		log.Fatalf("Failed to parse config file: #%v ", errUnmarshal)
	}

	return c
}
