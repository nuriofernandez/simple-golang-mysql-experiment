package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfig() *Settings {
	// Create a new Viper configuration parser instance
	config := viper.New()

	// Set the file name of the configurations file
	config.SetConfigName("database-settings")
	config.SetConfigType("yml")

	// Set the path to look for the configurations file
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	settings := &Settings{}
	unmarshalErr := config.Unmarshal(settings)
	if unmarshalErr != nil {
		fmt.Printf("unable to decode into config struct, %v", unmarshalErr)
	}

	return settings
}
