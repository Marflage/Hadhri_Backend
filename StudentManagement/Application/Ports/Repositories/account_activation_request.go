package repositories

import (
	"context"
	accountactivationrequest "hadhri/StudentManagement/Domain/AccountActivationRequest"
)

type IAccountActivationRequest interface {
	Store(ctx context.Context, e accountactivationrequest.AccountActivationRequest) (*int, error)
}
