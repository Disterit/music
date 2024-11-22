package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"music/logger"
	"music/pkg/handler"
	"music/pkg/repository"
	"music/pkg/service"
	"music/serverAPI"
	"os"
)

func main() {
	err := initConfig()
	if err != nil {
		logger.Log.Error("error to read config", err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		logger.Log.Error("Error loading .env file")
	}

	db := repository.Connection(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Database: viper.GetString("db.database"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(serverAPI.ServerApi)

	err = srv.Run(handlers.InitRoutes(), viper.GetString("port"))
	if err != nil {
		logger.Log.Error("error to start server")
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
