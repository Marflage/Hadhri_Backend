package repositories

import (
	"context"
	accountactivationrequest "hadhri/StudentManagement/Domain/AccountActivationRequest"
)

type IAccountActivationRequest interface {
	Save(ctx context.Context, e accountactivationrequest.AccountActivationRequest) (*int, error)
	Exists(ctx context.Context, id uint) (*bool, error)
	Approve(ctx context.Context, id uint) error
}
