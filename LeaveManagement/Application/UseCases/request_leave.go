package usecases

import (
	"context"
	commands "hadhri/LeaveManagement/Application/Commands"
	repositories "hadhri/LeaveManagement/Application/Ports/Repositories"
	domain "hadhri/LeaveManagement/Domain/LeaveRequest"
)

type RequestLeave struct {
	repo repositories.ILeaveRequest
}

func NewRequestLeaveUseCase(repo repositories.ILeaveRequest) RequestLeave {
	return RequestLeave{repo: repo}
}

func (uc RequestLeave) Execute(ctx context.Context, cmd commands.RequestLeave) error {
	entityPtr, err := domain.NewLeaveRequest(cmd.StudentId, cmd.StartDate, cmd.EndDate, cmd.Reason)

	if err != nil {
		return err
	}

	if err := uc.repo.AlreadyExists(ctx, *entityPtr); err != nil {
		return err
	}

	return uc.repo.AddLeaveRequest(ctx, *entityPtr)
}
