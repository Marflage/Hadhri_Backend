package ports

import (
	"context"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"
)

type IClassSessionRepo interface {
	Create(ctx context.Context, classSession dbmodels.ClassSession) error
	GetAll(ctx context.Context) ([]dbmodels.ClassSession, error)
}
