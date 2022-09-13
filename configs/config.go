package configs

import (
	"fmt"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var viper_set *viper.Viper

func Initialize() {

}

func init() {
	viper_set = viper.New()

	viper_set.SetConfigName(".env") //Name fof the file
	viper_set.SetConfigType("env")  // tye of file
	viper_set.AddConfigPath(".") // File location

	err := viper_set.ReadInConfig()

	if err != nil { 
		panic(fmt.Errorf("error when reading config file: %w", err))
	}

	viper_set.AutomaticEnv()
}

func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return Get(envName)
}

func Get(path string, defaultValue ...interface{}) interface{} {
	if !viper_set.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper_set.Get(path)
}

func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}