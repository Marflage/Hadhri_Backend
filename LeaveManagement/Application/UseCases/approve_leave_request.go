package usecases

import (
	"context"
	"errors"
	repositories "hadhri/LeaveManagement/Application/Ports/Repositories"
)

type ApproveLeaveRequest struct {
	repo repositories.ILeaveRequest
}

func NewApproveLeaveRequestUseCase(repo repositories.ILeaveRequest) ApproveLeaveRequest {
	return ApproveLeaveRequest{repo: repo}
}

func (self ApproveLeaveRequest) Execute(ctx context.Context, id uint) error {
	exists, err := self.repo.Exists(ctx, id)

	if err != nil {
		return err
	}

	if !*exists {
		return errors.New("Leave request does not exist.")
	}

	if err := self.repo.Approve(ctx, id); err != nil {
		return err
	}

	return nil
}
