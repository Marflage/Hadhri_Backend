package ports

import (
	"context"
	entities "hadhri/Admin/Domain/Entities"
	querymodels "hadhri/Admin/Domain/QueryModels"
)

type ICoursePlanRepo interface {
	Create(ctx context.Context, entity entities.CoursePlan) error
	GetAll(ctx context.Context) ([]querymodels.CoursePlan, error)
}
