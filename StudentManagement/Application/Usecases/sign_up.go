package usecases

import (
	"context"
	"fmt"
	commands "hadhri/StudentManagement/Application/Commands"
	queryservices "hadhri/StudentManagement/Application/Ports/QueryServices"
	repositories "hadhri/StudentManagement/Application/Ports/Repositories"
	services "hadhri/StudentManagement/Application/Ports/Services"
	student "hadhri/StudentManagement/Domain/Student"

	"golang.org/x/crypto/bcrypt"
)

type SignUp struct {
	coursePlanQueryService queryservices.ICoursePlan
	studentRepo            repositories.IStudent
	tokenService           services.TokenService
}

func NewSignUpUseCase(
	qs queryservices.ICoursePlan,
	repo repositories.IStudent,
	tokenSerivce services.TokenService,
) SignUp {
	return SignUp{
		coursePlanQueryService: qs,
		studentRepo:            repo,
		tokenService:           tokenSerivce,
	}
}

func (uc SignUp) Execute(ctx context.Context, cmd commands.SignUp) (*string, error) {
	coursePlanId, err := uc.coursePlanQueryService.GetId(ctx, cmd.CourseId, cmd.ClassScheduleId, cmd.ClassSessionId, true, cmd.Semester)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve course plan ID: %w", err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)

	if err != nil {
		// TODO: Log.
		return nil, fmt.Errorf("Failed to hash password: %w", err)
	}

	studentPtr, err := student.NewStudent(cmd.FirstName, cmd.LastName, cmd.Email, cmd.PhoneNumber, string(passwordHash), coursePlanId, cmd.Semester)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new student: %w", err)
	}

	studentIdPtr, err := uc.studentRepo.SignUp(ctx, *studentPtr)

	if err != nil {
		// TODO: Transform infra errors into domain ones by checking error codes.
		return nil, fmt.Errorf("Sign up failed: %w", err)
	}

	token, err := uc.tokenService.GenerateToken(*studentIdPtr)

	if err != nil {
		// TODO: Log.
		return nil, fmt.Errorf("Failed to generate token: %w", err)
	}

	return &token, nil
}
