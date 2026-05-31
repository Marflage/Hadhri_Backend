package usecases

import (
	"context"
	"errors"
	repositories "hadhri/LeaveManagement/Application/Ports/Repositories"
)

type CancelLeaveRequest struct {
	repo repositories.ILeaveRequest
}

func NewCancelLeaveRequestUseCase(repo repositories.ILeaveRequest) CancelLeaveRequest {
	return CancelLeaveRequest{repo: repo}
}

func (self CancelLeaveRequest) Execute(ctx context.Context, id uint, studentId uint) error {
	exists, err := self.repo.Exists(ctx, id, studentId)

	if err != nil {
		// TODO: Might need to return domain error.
		return err
	}

	if !*exists {
		return errors.New("Leave request does not exist.")
	}

	if err := self.repo.Cancel(ctx, id, studentId); err != nil {
		// TODO: Might need to return domain error.
		return err
	}

	return nil
}
