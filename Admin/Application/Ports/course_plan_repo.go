package ports

import (
	"context"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"
)

type ICoursePlanRepo interface {
	Create(ctx context.Context, entity dbmodels.CoursePlan) error
}
