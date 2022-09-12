package main

import (
	"go-gin-crud/handlers"
	"go-gin-crud/services"

	
)

func main()  {
	services := services.NewServices()
	router := handlers.NewRouter(services)

	router.Run(":5000")
}