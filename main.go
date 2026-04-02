package main

import (
	handlers "hadhri/Handlers"
	middleware "hadhri/Middleware"
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

	router.GET("/students", middleware.AuthMiddleware(), handlers.GetStudentDetails)
	router.GET("/student-enrollments", middleware.AuthMiddleware(), handlers.GetStudentEnrollmentDetails)
	// TODO: Move this route behind the IP-whitelisting middleware.
	router.POST("/log-attendance", middleware.AuthMiddleware(), handlers.LogAttendance)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
