package main

import (
	"context"
	adminUsecases "hadhri/Admin/Application/UseCases"
	infrastructure "hadhri/Admin/Infrastructure"
	adminQueryservices "hadhri/Admin/Infrastructure/QueryServices"
	adminWebapi "hadhri/Admin/WebApi"
	handlers "hadhri/Handlers"
	middleware "hadhri/Middleware"
	usecases "hadhri/StudentManagement/Application/Usecases"
	queryservices "hadhri/StudentManagement/Infrastructure/QueryServices"
	repositories "hadhri/StudentManagement/Infrastructure/Repositories"
	services "hadhri/StudentManagement/Infrastructure/Services"
	webapi "hadhri/StudentManagement/WebApi"
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
	addCourseUC := adminUsecases.NewAddCourseUseCase(courseRepo)
	getAllCoursesUC := adminUsecases.NewGetAllCoursesUseCase(courseRepo)
	addCourseHandler := adminWebapi.NewAddCourseHandler(addCourseUC)
	getAllCoursesHandler := adminWebapi.NewGetAllCoursesHandler(getAllCoursesUC)

	classScheduleRepo := infrastructure.NewClassScheduleRepo(pool)
	addClassScheduleUC := adminUsecases.NewAddClassScheduleUseCase(classScheduleRepo)
	getAllClassSchedulesUC := adminUsecases.NewGetAllClassSchedulesUseCase(classScheduleRepo)
	addClassScheduleHandler := adminWebapi.NewClassScheduleHandler(addClassScheduleUC)
	getAllClassSchedulesHandler := adminWebapi.NewGetAllClassSchedulesHandler(getAllClassSchedulesUC)

	classSessionRepo := infrastructure.NewClassSessionRepo(pool)
	addClassSessionUC := adminUsecases.NewAddClassSessionUseCase(classSessionRepo)
	getAllClassSessionsUC := adminUsecases.NewGetAllClassSessionsUseCase(classSessionRepo)
	addClassSessionHandler := adminWebapi.NewAddClassSessionHandler(addClassSessionUC)
	getAllClassSessionsHandler := adminWebapi.NewGetAllClassSessionsHandler(getAllClassSessionsUC)

	coursePlanRepo := infrastructure.NewCoursePlanRepo(pool)
	coursePlanQueryService := adminQueryservices.NewCoursePlanQueryService(pool)
	addCoursePlanUC := adminUsecases.NewAddCoursePlanUseCase(coursePlanRepo)
	getAllCoursePlansUC := adminUsecases.NewGetAllCoursePlansUseCase(coursePlanQueryService)
	addCoursePlanHandler := adminWebapi.NewAddCoursePlanHandler(addCoursePlanUC)
	getAllCoursePlansHandler := adminWebapi.NewGetAllCoursePlansHandler(getAllCoursePlansUC)

	studentRepo := repositories.NewStudentRepo(pool)
	studentQueryService := adminQueryservices.NewStudentQueryService(pool)
	addStudentUC := usecases.NewAddStudentUseCase(studentRepo)
	addStudentHandler := webapi.NewAddStudentHandler(addStudentUC)
	getStudentUC := adminUsecases.NewGetStudentUseCase(studentQueryService)
	getStudentHandler := adminWebapi.NewGetStudentHandler(getStudentUC)

	// TODO: Rename for better organization.
	coursePlanQueryService2 := queryservices.NewCoursePlanQueryService(pool)
	studentRepo2 := repositories.NewStudentRepo(pool)
	tokenService := services.NewJwtService()
	signUpUsecase := usecases.NewSignUpUseCase(coursePlanQueryService2, studentRepo2, tokenService)
	signUpHandler := webapi.NewSignUpHandler(signUpUsecase)

	r := gin.Default()

	// Admin endpoints
	r.POST("/course", addCourseHandler.AddCourse)
	r.POST("/class-schedule", addClassScheduleHandler.AddClassSchedule)
	r.POST("/class-session", addClassSessionHandler.AddClassSession)
	r.POST("/course-plan", addCoursePlanHandler.AddCoursePlan)

	// TODO: Should the path be plural?
	r.GET("/courses", getAllCoursesHandler.GetAll)
	r.GET("/class-schedules", getAllClassSchedulesHandler.GetAll)
	r.GET("/class-sessions", getAllClassSessionsHandler.GetAll)
	r.GET("/admin/course-plans", getAllCoursePlansHandler.GetAll)

	r.POST("/student", addStudentHandler.Handle)
	r.GET("/student", getStudentHandler.Handle)

	// TODO: Create a middleware to handle exceptions.
	// TODO: Create a middleware to format errors and send them in response.
	r.POST("/sign-up", signUpHandler.Handle)
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
