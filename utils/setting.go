package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Mode     string
	Database struct {
		Name     string
		Host     string
		Port     string
		Username string
		Password string
	}
	Server struct {
		Host string
		Port string
	}
}

var C Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("err config: %s \n", err)
	}
	viper.Unmarshal(&C)
	fmt.Println(C)
}
