package main

import (
	"blog/database"
	"blog/routes"
)

func main() {
	database.Connect()

	// Initialize routes
	router := routes.SetupRoutes()

	// Start the server
	router.Listen(":8000")
}
