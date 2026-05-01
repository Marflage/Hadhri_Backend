package ports

import (
	"context"
	entities "hadhri/Admin/Domain/Entities"
)

type ICourseRepo interface {
	Create(ctx context.Context, course entities.Course) error
	GetAll(ctx context.Context) ([]entities.Course, error)
}
