package usecases

import (
	"context"
	"errors"
	repositories "hadhri/LeaveManagement/Application/Ports/Repositories"
)

type RejectLeaveRequest struct {
	repo repositories.ILeaveRequest
}

func NewRejectLeaveRequestUseCase(repo repositories.ILeaveRequest) RejectLeaveRequest {
	return RejectLeaveRequest{repo: repo}
}

func (self RejectLeaveRequest) Execute(ctx context.Context, id uint) error {
	exists, err := self.repo.Exists(ctx, id)

	if err != nil {
		return err
	}

	if !*exists {
		return errors.New("Leave request does not exist.")
	}

	return self.repo.Reject(ctx, id)
}
