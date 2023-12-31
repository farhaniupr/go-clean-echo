package library

import (
	"log"

	"github.com/spf13/viper"
)

// Variabel env
type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENV"`
	LogOutput   string `mapstructure:"LOG_OUTPUT"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	DBUsername  string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASS"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBName      string `mapstructure:"DB_NAME"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
}

// Get environment
func ModuleEnv() Env {

	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Println("environment can't be loaded: ", err.Error())
	}

	return env
}
