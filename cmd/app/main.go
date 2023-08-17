package main

import (
	"awesomeGoProject/cmd/app/migration"
	"awesomeGoProject/internal/handler"
	"awesomeGoProject/internal/repository"
	"awesomeGoProject/internal/server"
	"awesomeGoProject/internal/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}

	if err := migration.InitMigrations(); err != nil {
		log.Fatalf("Error managing database migrations: %s", err.Error())
	}

	repo := repository.NewRepository()
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
