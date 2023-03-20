package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	DBname     string
	DbPort     string
	DbDriver   string
	DbPassword string
)

func init() {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error Membaca File, %s", err.Error())
	}
	DBname = viper.GetString("db.name")
	DbPort = viper.GetString("db.port")
	DbDriver = viper.GetString("db.driver")
	DbPassword = viper.GetString("db.password")
}
