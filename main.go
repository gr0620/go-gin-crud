package main

import (
	"github.com/gr0620/go-gin-crud/handlers"
	"github.com/gr0620/go-gin-crud/services"

	
)

func main()  {
	services := services.NewServices()
	router := handlers.NewRouter(services)

	router.Run(":5000")
}