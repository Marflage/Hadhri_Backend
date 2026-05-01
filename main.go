package main

import (
	"context"
	usecases "hadhri/Admin/Application/UseCases"
	infrastructure "hadhri/Admin/Infrastructure"
	webapi "hadhri/Admin/WebApi"
	handlers "hadhri/Handlers"
	middleware "hadhri/Middleware"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("noBlank", noBlankValidator)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	connStr := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(ctx, connStr)

	if err != nil {
		log.Fatalf("Failed to create DB pool: %v", err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	courseRepo := infrastructure.NewCourseRepo(pool)
	addCourseUC := usecases.NewAddCourseUseCase(courseRepo)
	getAllCoursesUC := usecases.NewGetAllCoursesUseCase(courseRepo)
	addCourseHandler := webapi.NewAddCourseHandler(addCourseUC)
	getAllCoursesHandler := webapi.NewGetAllCoursesHandler(getAllCoursesUC)

	classScheduleRepo := infrastructure.NewClassScheduleRepo(pool)
	addClassScheduleUC := usecases.NewAddClassScheduleUseCase(classScheduleRepo)
	getAllClassSchedulesUC := usecases.NewGetAllClassSchedulesUseCase(classScheduleRepo)
	addClassScheduleHandler := webapi.NewClassScheduleHandler(addClassScheduleUC)
	getAllClassSchedulesHandler := webapi.NewGetAllClassSchedulesHandler(getAllClassSchedulesUC)

	classSessionRepo := infrastructure.NewClassSessionRepo(pool)
	addClassSessionUC := usecases.NewAddClassSessionUseCase(classSessionRepo)
	addClassSessionHandler := webapi.NewAddClassSessionHandler(addClassSessionUC)

	coursePlanRepo := infrastructure.NewCoursePlanRepo(pool)
	addCoursePlanUC := usecases.NewAddCoursePlanUseCase(coursePlanRepo)
	addCoursePlanHandler := webapi.NewAddCoursePlanHandler(addCoursePlanUC)

	r := gin.Default()

	// Admin endpoints
	r.POST("/course", addCourseHandler.AddCourse)
	r.POST("/class-schedule", addClassScheduleHandler.AddClassSchedule)
	r.POST("/class-session", addClassSessionHandler.AddClassSession)
	r.POST("/course-plan", addCoursePlanHandler.AddCoursePlan)

	// TODO: Should the path be plural?
	r.GET("/courses", getAllCoursesHandler.GetAll)
	r.GET("/class-schedules", getAllClassSchedulesHandler.GetAll)

	// TODO: Create a middleware to handle exceptions.
	// TODO: Create a middleware to format errors and send them in response.
	r.POST("/sign-up", handlers.SignUp)
	r.POST("/sign-in", handlers.SignIn)

	// TODO: Create an authentication middleware.

	// Reference data routing
	// TODO: Add authentication for this route.
	r.GET("/course-plans", handlers.GetCoursePlans)

	r.GET("/students", middleware.AuthMiddleware(), handlers.GetStudentDetails)
	r.GET("/student-enrollments", middleware.AuthMiddleware(), handlers.GetStudentEnrollmentDetails)
	r.GET("/attendance-status", middleware.AuthMiddleware(), handlers.IsAttendanceLogged)
	// TODO: Move this route behind the IP-whitelisting middleware.
	r.POST("/log-attendance", middleware.AuthMiddleware(), handlers.LogAttendance)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
