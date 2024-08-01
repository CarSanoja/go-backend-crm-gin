package main

import (
	"go-backend-crm/config"
	"go-backend-crm/handlers"
	"log"
)

func main() {
	config.LoadConfig()
	if err := handlers.LoadCustomers(); err != nil {
		log.Fatalf("Failed to load customers: %v", err)
	}

	r := handlers.SetupRouter()
	log.Printf("Server running at port %s", config.GetConfig().Port)
	if err := r.Run(":" + config.GetConfig().Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
