package ports

import (
	"context"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"
)

type ICourseRepo interface {
	Create(ctx context.Context, course dbmodels.Course) error
	GetAll(ctx context.Context) ([]dbmodels.Course, error)
}
