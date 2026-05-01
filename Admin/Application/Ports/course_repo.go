package ports

import (
	"context"
	domain "hadhri/Admin/Domain"
)

type ICourseRepo interface {
	Create(ctx context.Context, course domain.Course) error
	GetAll(ctx context.Context) ([]domain.Course, error)
}
