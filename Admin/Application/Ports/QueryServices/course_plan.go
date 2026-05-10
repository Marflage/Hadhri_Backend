package queryservices

import (
	"context"
	dtos "hadhri/Admin/Application/Dtos"
)

type ICoursePlan interface {
	GetAll(ctx context.Context) ([]dtos.CoursePlan, error)
}
