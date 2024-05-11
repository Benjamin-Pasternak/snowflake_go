package util

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// TODO:
// - write a function that finds the defaults, checks if there is a value there in the environment...
// 	if there is set that as the default... if not use the default value ... else do not set.

// psudo code:
// func read_config() {
// 	read file into memory
// 	create regex to find all instances of these variables and their defaults (potentially)
// 	find instances
// 	for each instance:
// 		if variable is in environment
// 			continue
// 		else if variable is not in environment and there is a default
// 			set this value in the config
// 		else
// 			do nothing
// 	return config
// }

// let us keep it to just one profile for now ... don't really think I need more than one but incase I wanna reuse in other projects ... could be useful
func InitConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msg("Failed to read application yaml")
		os.Exit(1)
	}

	// automatically maps the environment variables that match the keys in yaml (case insensitive)
	// e.g., if env PORT=8080 and in yaml port: 8081 it should take 8080
	viper.AutomaticEnv()
}
