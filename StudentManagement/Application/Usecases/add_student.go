package usecases

import (
	"context"
	"crypto/rand"
	"fmt"
	commands "hadhri/StudentManagement/Application/Commands"
	dtos "hadhri/StudentManagement/Application/Dtos"
	queryservices "hadhri/StudentManagement/Application/Ports/QueryServices"
	repositories "hadhri/StudentManagement/Application/Ports/Repositories"
	student "hadhri/StudentManagement/Domain/Student"
	"math/big"
)

type AddStudent struct {
	studentRepo            repositories.IStudent
	coursePlanQueryService queryservices.ICoursePlan
}

func NewAddStudentUseCase(repo repositories.IStudent, queryService queryservices.ICoursePlan) AddStudent {
	return AddStudent{studentRepo: repo, coursePlanQueryService: queryService}
}

func (uc AddStudent) Execute(ctx context.Context, cmd commands.AddStudent) (*dtos.StudentCredentials, error) {
	coursePlanId, err := uc.coursePlanQueryService.GetId(ctx, cmd.CourseId, cmd.ClassScheduleId, cmd.ClassSessionId, true, cmd.Semester)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve course plan ID: %w", err)
	}

	tempPassword, err := generateTempPassword()

	if err != nil {
		return nil, err
	}

	studentPtr, err := student.NewStudent(cmd.FirstName, cmd.LastName, cmd.Email, cmd.PhoneNumber, string(tempPassword), coursePlanId, cmd.Semester)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new student: %w", err)
	}

	if _, err := uc.studentRepo.SignUp(ctx, *studentPtr); err != nil {
		// TODO: Transform infra errors into domain ones by checking error codes.
		return nil, fmt.Errorf("Sign up failed: %w", err)
	}

	creds := dtos.StudentCredentials{
		Email:    cmd.Email,
		Password: tempPassword,
	}

	return &creds, nil
}

const tempPasswordLength = 8

var passwordCharset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*")

func generateTempPassword() (string, error) {
	b := make([]byte, tempPasswordLength)
	for i := range b {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(passwordCharset))))
		if err != nil {
			return "", err
		}
		b[i] = passwordCharset[idx.Int64()]
	}
	return string(b), nil
}
