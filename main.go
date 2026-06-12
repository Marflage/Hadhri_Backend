package main

import (
	"context"
	adminUsecases "hadhri/Admin/Application/UseCases"
	adminInfrastructure "hadhri/Admin/Infrastructure"
	adminQueryservices "hadhri/Admin/Infrastructure/QueryServices"
	adminWebapi "hadhri/Admin/WebApi"
	orphanHandlers "hadhri/Handlers"
	lvMgtUseCases "hadhri/LeaveManagement/Application/UseCases"
	lvMgtRepositories "hadhri/LeaveManagement/Infrastructure/Repositories"
	lvMgtHandlers "hadhri/LeaveManagement/WebApi/Handlers"
	middleware "hadhri/Middleware"
	stdntMgtUseCases "hadhri/StudentManagement/Application/Usecases"
	stdntMgtQueryServices "hadhri/StudentManagement/Infrastructure/QueryServices"
	stdntMgtRepositories "hadhri/StudentManagement/Infrastructure/Repositories"
	stdntMgtServices "hadhri/StudentManagement/Infrastructure/Services"
	stdntMgtHandlers "hadhri/StudentManagement/WebApi/Handlers"
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

	courseRepo := adminInfrastructure.NewCourseRepo(pool)
	addCourseUC := adminUsecases.NewAddCourseUseCase(courseRepo)
	getAllCoursesUC := adminUsecases.NewGetAllCoursesUseCase(courseRepo)
	addCourseHandler := adminWebapi.NewAddCourseHandler(addCourseUC)
	getAllCoursesHandler := adminWebapi.NewGetAllCoursesHandler(getAllCoursesUC)

	classScheduleRepo := adminInfrastructure.NewClassScheduleRepo(pool)
	addClassScheduleUC := adminUsecases.NewAddClassScheduleUseCase(classScheduleRepo)
	getAllClassSchedulesUC := adminUsecases.NewGetAllClassSchedulesUseCase(classScheduleRepo)
	addClassScheduleHandler := adminWebapi.NewClassScheduleHandler(addClassScheduleUC)
	getAllClassSchedulesHandler := adminWebapi.NewGetAllClassSchedulesHandler(getAllClassSchedulesUC)

	classSessionRepo := adminInfrastructure.NewClassSessionRepo(pool)
	addClassSessionUC := adminUsecases.NewAddClassSessionUseCase(classSessionRepo)
	getAllClassSessionsUC := adminUsecases.NewGetAllClassSessionsUseCase(classSessionRepo)
	addClassSessionHandler := adminWebapi.NewAddClassSessionHandler(addClassSessionUC)
	getAllClassSessionsHandler := adminWebapi.NewGetAllClassSessionsHandler(getAllClassSessionsUC)

	coursePlanRepo := adminInfrastructure.NewCoursePlanRepo(pool)
	coursePlanQueryService := adminQueryservices.NewCoursePlanQueryService(pool)
	addCoursePlanUC := adminUsecases.NewAddCoursePlanUseCase(coursePlanRepo)
	getAllCoursePlansUC := adminUsecases.NewGetAllCoursePlansUseCase(coursePlanQueryService)
	addCoursePlanHandler := adminWebapi.NewAddCoursePlanHandler(addCoursePlanUC)
	getAllCoursePlansHandler := adminWebapi.NewGetAllCoursePlansHandler(getAllCoursePlansUC)

	studentQueryService := adminQueryservices.NewStudentQueryService(pool)

	getStudentUC := adminUsecases.NewGetStudentUseCase(studentQueryService)
	getStudentHandler := adminWebapi.NewGetStudentHandler(getStudentUC)

	// TODO: Rename for better organization.

	// Student Management
	tokenService := stdntMgtServices.NewJwtService()

	studentRepo := stdntMgtRepositories.NewStudentRepo(pool)
	accountActivationRequestRepo := stdntMgtRepositories.NewAccountActivationRequestRepo(pool)

	coursePlanQueryService2 := stdntMgtQueryServices.NewCoursePlanQueryService(pool)

	signUpUseCase := stdntMgtUseCases.NewSignUpUseCase(coursePlanQueryService2, accountActivationRequestRepo, tokenService)
	signUpHandler := stdntMgtHandlers.NewSignUpHandler(signUpUseCase)

	addStudentUC := stdntMgtUseCases.NewAddStudentUseCase(studentRepo, coursePlanQueryService2)
	addStudentHandler := stdntMgtHandlers.NewAddStudentHandler(addStudentUC)

	approveAccountActivationRequestUseCase := stdntMgtUseCases.NewApproveAccountActivationUseCase(accountActivationRequestRepo)
	approveAccountActivationRequestHandler := stdntMgtHandlers.NewApproveAccountActivationHandler(approveAccountActivationRequestUseCase)

	declineAccountActivationRequestUseCase := stdntMgtUseCases.NewDeclineAccountActivationUseCase(accountActivationRequestRepo)
	declineAccountActivationRequestHandler := stdntMgtHandlers.NewDeclineAccountActivationHandler(declineAccountActivationRequestUseCase)

	// Leave Management
	leaveRequestRepo := lvMgtRepositories.NewLeaveRequestRepo(pool)

	requestLeaveUseCase := lvMgtUseCases.NewRequestLeaveUseCase(leaveRequestRepo)
	requestLeaveHandler := lvMgtHandlers.NewRequestLeaveHandler(requestLeaveUseCase)

	editLeaveRequestUseCase := lvMgtUseCases.NewEditLeaveRequestUseCase(leaveRequestRepo)
	editLeaveRequestHandler := lvMgtHandlers.NewEditLeaveHandler(editLeaveRequestUseCase)

	cancelLeaveRequestUseCase := lvMgtUseCases.NewCancelLeaveRequestUseCase(leaveRequestRepo)
	cancelLeaveRequestHandler := lvMgtHandlers.NewCancelLeaveRequestHandler(cancelLeaveRequestUseCase)

	approveLeaveRequestUseCase := lvMgtUseCases.NewApproveLeaveRequestUseCase(leaveRequestRepo)
	approveLeaveRequestHandler := lvMgtHandlers.NewApproveLeaveRequestHandler(approveLeaveRequestUseCase)

	rejectLeaveRequestUseCase := lvMgtUseCases.NewRejectLeaveRequestUseCase(leaveRequestRepo)
	rejectLeaveRequestHandler := lvMgtHandlers.NewRejectLeaveRequest(rejectLeaveRequestUseCase)

	r := gin.Default()

	// TODO: Add authorization to all the endpoints except for the sign-up and sign-in ones.

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

	r.POST("/leaves", requestLeaveHandler.Handle)
	r.PATCH("/leaves/:id", editLeaveRequestHandler.Handle)
	r.PATCH("/leaves/:id/cancel", cancelLeaveRequestHandler.Handle)
	r.PATCH("/leaves/:id/approve", approveLeaveRequestHandler.Handle)
	r.PATCH("/leaves/:id/reject", rejectLeaveRequestHandler.Handle)

	r.POST("/sign-up", signUpHandler.Handle)
	r.PATCH("/account-activation/:id/approve", approveAccountActivationRequestHandler.Handle)
	r.PATCH("/account-activation/:id/decline", declineAccountActivationRequestHandler.Handle)

	// TODO: Create a middleware to handle exceptions.
	// TODO: Create a middleware to format errors and send them in response.
	r.POST("/sign-in", orphanHandlers.SignIn)

	// TODO: Create an authentication middleware.

	// Reference data routing
	// TODO: Add authentication for this route.
	r.GET("/course-plans", orphanHandlers.GetCoursePlans)

	r.GET("/students", middleware.AuthMiddleware(), orphanHandlers.GetStudentDetails)
	r.GET("/student-enrollments", middleware.AuthMiddleware(), orphanHandlers.GetStudentEnrollmentDetails)
	r.GET("/attendance-status", middleware.AuthMiddleware(), orphanHandlers.IsAttendanceLogged)
	// TODO: Move this route behind the IP-whitelisting middleware.
	r.POST("/log-attendance", middleware.AuthMiddleware(), orphanHandlers.LogAttendance)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
