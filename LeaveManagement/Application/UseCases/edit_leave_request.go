package usecases

import (
	"context"
	"fmt"
	commands "hadhri/LeaveManagement/Application/Commands"
	repositories "hadhri/LeaveManagement/Application/Ports/Repositories"
)

type EditLeaveRequest struct {
	repo repositories.ILeaveRequest
}

func NewEditLeaveRequestUseCase(repo repositories.ILeaveRequest) EditLeaveRequest {
	return EditLeaveRequest{repo: repo}
}

func (self EditLeaveRequest) Execute(ctx context.Context, cmd commands.EditLeaveRequest) error {
	e, err := self.repo.Get(ctx, cmd.Id)

	if err != nil {
		return fmt.Errorf("Failed to retrieve leave request.")
	}

	if cmd.StartDate != nil || cmd.EndDate != nil {
		if err := e.Reschedule(cmd.StartDate, cmd.EndDate); err != nil {
			return err
		}

		if err := self.repo.AlreadyExists(ctx, *e); err != nil {
			return err
		}
	}

	if cmd.Reason != nil {
		if err := e.ChangeReason(*cmd.Reason); err != nil {
			return err
		}
	}

	if e.GetStatus() != "pending" {
		return fmt.Errorf("Leave request cannot be edited.")
	}

	if err := self.repo.Update(ctx, *e); err != nil {
		return err
	}

	return nil
}
