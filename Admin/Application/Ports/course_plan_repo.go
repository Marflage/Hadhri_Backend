package ports

import (
	"context"
	querymodels "hadhri/Admin/Domain/QueryModels"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"
)

type ICoursePlanRepo interface {
	Create(ctx context.Context, entity dbmodels.CoursePlan) error
	GetAll(ctx context.Context) ([]querymodels.CoursePlan, error)
}
