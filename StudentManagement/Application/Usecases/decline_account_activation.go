package usecases

import (
	"context"
	"errors"
	"fmt"
	repositories "hadhri/StudentManagement/Application/Ports/Repositories"
)

type DeclineAccountActivation struct {
	repo repositories.IAccountActivationRequest
}

func NewDeclineAccountActivationUseCase(repo repositories.IAccountActivationRequest) DeclineAccountActivation {
	return DeclineAccountActivation{repo: repo}
}

func (self DeclineAccountActivation) Execute(ctx context.Context, id uint) error {
	exists, err := self.repo.Exists(ctx, id)

	if err != nil {
		return fmt.Errorf("Failed to decline account activation request: %w", err)
	}

	if !*exists {
		return errors.New("Account activation request does not exist.")
	}

	return self.repo.Decline(ctx, id)
}
