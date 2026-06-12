package usecases

import (
	"context"
	"errors"
	"fmt"
	repositories "hadhri/StudentManagement/Application/Ports/Repositories"
)

type ApproveAccountActivation struct {
	repo repositories.IAccountActivationRequest
}

func NewApproveAccountActivationUseCase(repo repositories.IAccountActivationRequest) ApproveAccountActivation {
	return ApproveAccountActivation{repo: repo}
}

func (self ApproveAccountActivation) Execute(ctx context.Context, id uint) error {
	exists, err := self.repo.Exists(ctx, id)

	if err != nil {
		return fmt.Errorf("Failed to approve activation request: %w", err)
	}

	if !*exists {
		return errors.New("Account activation request does not exist.")
	}

	return self.repo.Approve(ctx, id)
}
