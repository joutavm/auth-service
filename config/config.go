package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init() *viper.Viper {

	env := os.Getenv("SCOPE")
	if env == "" {
		env = "local"
	}


	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("config/resource/")
	viper.AutomaticEnv()

	err := config.ReadInConfig()

	for _, k := range config.AllKeys() {
		value := config.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			config.Set(k, os.Getenv(strings.TrimSuffix(strings.TrimPrefix(value,"${"), "}")))
		}
	}

	if err != nil {
		log.Fatal("error on parsing configuration file", err)
	}
	return config
}

