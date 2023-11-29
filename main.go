package main

import (
	"project-pertama-golang/routes"
)

func main() {
	router := routes.SetupRouter()

	// Start the Gin server
	router.Run(":666")
}
