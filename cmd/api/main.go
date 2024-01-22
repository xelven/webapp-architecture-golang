package main

import (
	"fmt"
	"net/http"

	"webapp-core/config"
	"webapp-core/core/routers"
	"webapp-core/util/logger"
)

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	serverConfig := config.New()
	serverLogger := logger.New(serverConfig.Server.Debug)

	dbString := fmt.Sprintf(fmtDBString,
		serverConfig.DB.Host,
		serverConfig.DB.Username,
		serverConfig.DB.Password,
		serverConfig.DB.DBName,
		serverConfig.DB.Port,
	)
	// go to db connection
	fmt.Println("connect db with ", dbString)

	routers := routers.New(serverLogger)

	// create server
	apiServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", serverConfig.Server.Port),
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
