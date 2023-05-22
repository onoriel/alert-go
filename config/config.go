package config

import (
	"log"

	"github.com/spf13/viper"
)

// Database attributes
type Config struct {
	Server   string
	Database string
}

// Read configuration file
func (c *Config) Read() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
		return
	}
	c.Server = viper.GetString("database.url")
	c.Database = viper.GetString("database.name")
}
