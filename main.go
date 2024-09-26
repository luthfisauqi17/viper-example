package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	appName := v.GetString("name")
	appPort := v.GetInt("port")
	dbDriver := v.GetString("database.driver")
	dbHost := v.GetString("database.host")
	dbPort := v.GetInt("database.port")

	fmt.Printf("Application Name: %s\n", appName)
	fmt.Printf("Application Port: %d\n", appPort)
	fmt.Printf("Database Driver: %s\n", dbDriver)
	fmt.Printf("Database Host: %s\n", dbHost)
	fmt.Printf("Database Port: %d\n", dbPort)
}
