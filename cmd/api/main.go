package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

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
	var dbLogLevel gormlogger.LogLevel
	if serverConfig.DB.Debug {
		dbLogLevel = gormlogger.Info
	} else {
		dbLogLevel = gormlogger.Error
	}

	masterDataSource, err := gorm.Open(postgres.Open(dbString), &gorm.Config{Logger: gormlogger.Default.LogMode(dbLogLevel)})
	if err != nil {
		serverLogger.Fatal().Err(err).Msg("DB connection start failure")
		return
	}

	routers := routers.New(serverLogger, masterDataSource)

	// create server
	apiServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", serverConfig.Server.Port),
		Handler:      routers,
		ReadTimeout:  serverConfig.Server.TimeoutRead,
		WriteTimeout: serverConfig.Server.TimeoutWrite,
		IdleTimeout:  serverConfig.Server.TimeoutIdle,
	}

	// start server
	serverLogger.Info().Msgf("Starting server %v", apiServer.Addr)
	if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		serverLogger.Fatal().Err(err).Msg("Server startup failure")
	}

}
