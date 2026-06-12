package usecases

import (
	"context"
	"fmt"
	commands "hadhri/StudentManagement/Application/Commands"
	queryservices "hadhri/StudentManagement/Application/Ports/QueryServices"
	repositories "hadhri/StudentManagement/Application/Ports/Repositories"
	services "hadhri/StudentManagement/Application/Ports/Services"
	accountactivationrequest "hadhri/StudentManagement/Domain/AccountActivationRequest"
)

type SignUp struct {
	coursePlanQueryService       queryservices.ICoursePlan
	accountActivationRequestRepo repositories.IAccountActivationRequest
	tokenService                 services.TokenService
}

func NewSignUpUseCase(
	qs queryservices.ICoursePlan,
	repo repositories.IAccountActivationRequest,
	tokenSerivce services.TokenService,
) SignUp {
	return SignUp{
		coursePlanQueryService:       qs,
		accountActivationRequestRepo: repo,
		tokenService:                 tokenSerivce,
	}
}

func (self SignUp) Execute(ctx context.Context, cmd commands.SignUp) (*string, error) {
	coursePlanId, err := self.coursePlanQueryService.GetId(ctx, cmd.CourseId, cmd.ClassScheduleId, cmd.ClassSessionId, true, cmd.Semester)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve course plan ID: %w", err)
	}

	e, err := accountactivationrequest.NewAccountActivationRequest(cmd.FullName, cmd.Email, cmd.PhoneNumber, cmd.Password, coursePlanId, cmd.Semester)

	if err != nil {
		// TODO: Errors in domain-firendly-language should be thrown but the raw ones must be logged at the point of error.
		return nil, fmt.Errorf("Failed to request account activation: %w", err)
	}

	id, err := self.accountActivationRequestRepo.Save(ctx, *e)

	if err != nil {
		// TODO: Transform infra errors into domain ones by checking error codes.
		return nil, fmt.Errorf("Sign up failed: %w", err)
	}

	token, err := self.tokenService.GenerateToken(*id)

	if err != nil {
		// TODO: Log.
		return nil, fmt.Errorf("Failed to generate token: %w", err)
	}

	return &token, nil
}
