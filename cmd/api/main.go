package main

import (
	"fmt"
	"webapp-core/config"
)

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	serverConfig := config.New()
	dbString := fmt.Sprintf(fmtDBString, serverConfig.DB.Host, serverConfig.DB.Username, serverConfig.DB.Password, serverConfig.DB.DBName, serverConfig.DB.Port)
	// go to db connection
}
