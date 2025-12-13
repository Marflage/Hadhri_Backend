package main

import (
	handlers "hadhri/Handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// TODO: Create a middleware to handle exceptions.
	// TODO: Create a middleware to format errors and send them in response.
	router.POST("/sign-up", handlers.SignUp)
	router.POST("/sign-in", handlers.SignIn)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
