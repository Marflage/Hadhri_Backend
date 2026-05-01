package ports

import (
	"context"
	entities "hadhri/Admin/Domain/Entities"
)

type ICoursePlanRepo interface {
	Create(ctx context.Context, entity entities.CoursePlan) error
}
