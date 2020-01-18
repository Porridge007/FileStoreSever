package util

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func GetConfig(key string) interface{} {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file error: ", err)
		os.Exit(1)
	}
	toGet := viper.Get(key)
	return toGet
}
