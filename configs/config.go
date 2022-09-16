package configs

import (
	"log"

	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/spf13/viper"
)

var viper_set *viper.Viper

func Initialize() {

}

// Function called anytime we need to use setting on .env file
func LoadConfigEnv() entity.Config {
	var config entity.Config

	viper.SetConfigName(".env") //Name fof the file
	viper.SetConfigType("env")  // tye of file
	viper.AddConfigPath(".") // File location
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil { 
		log.Fatalf("Error while reading from .env: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil { 
		log.Fatalf("Error while unmarshalling .env file: %v", err)
	}

	return config
}
