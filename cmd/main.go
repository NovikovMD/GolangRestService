package main

import (
	"awesomeGoProject/internal/route"
	"awesomeGoProject/internal/server"
	"log"
)

func main() {
	handler := new(route.Initializer)
	srv := server.InitServer(handler.InitRoutes())
	if err := srv.Run(); err != nil {
		log.Fatalf("Error during server execution: %s", err.Error())
	}
}
