package ports

import (
	"context"
	domain "hadhri/Admin/Domain"
)

type ICoursePlanRepo interface {
	Create(ctx context.Context, entity domain.CoursePlan) error
}
