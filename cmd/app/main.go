package main

import (
	"awesomeGoProject/internal/handler"
	"awesomeGoProject/internal/repository"
	"awesomeGoProject/internal/server"
	"awesomeGoProject/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error initializing environment: %s", err.Error())
	}

	dbInstance, err := repository.NewPostgresInstance()
	if err != nil {
		log.Fatalf("Error managing initializing database: %s", err.Error())
	}

	repo := repository.NewRepository(dbInstance)
	serv := service.NewService(repo)
	hdr := handler.NewHandler(serv)

	srv := server.InitServer(viper.GetString("port"), hdr.InitRoutes())
	if err := srv.Run(); err != nil {
		log.Fatalf("Error during server execution: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
