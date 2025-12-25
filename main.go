package main

import (
	handlers "hadhri/Handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("noBlank", noBlankValidator)
	}
}

func main() {
	router := gin.Default()

	// TODO: Create a middleware to handle exceptions.
	// TODO: Create a middleware to format errors and send them in response.
	router.POST("/sign-up", handlers.SignUp)
	router.POST("/sign-in", handlers.SignIn)

	// TODO: Create an authentication middleware.

	// Reference data routing
	// TODO: Add authentication for this route.
	router.GET("/course-plans", handlers.GetCoursePlans)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
