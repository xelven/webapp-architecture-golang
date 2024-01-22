package main

import (
	"fmt"
	"net/http"

	"webapp-core/config"
	"webapp-core/core/routers"
)

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	serverConfig := config.New()
	dbString := fmt.Sprintf(fmtDBString,
		serverConfig.DB.Host,
		serverConfig.DB.Username,
		serverConfig.DB.Password,
		serverConfig.DB.DBName,
		serverConfig.DB.Port,
	)
	// go to db connection
	fmt.Println("connect db with ", dbString)

	routers := routers.New()

	// create server
	apiServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", 8080),
		Handler:      routers,
		ReadTimeout:  serverConfig.Server.TimeoutRead,
		WriteTimeout: serverConfig.Server.TimeoutWrite,
		IdleTimeout:  serverConfig.Server.TimeoutIdle,
	}

	// start server
	if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Server startup failure")
	}

}
