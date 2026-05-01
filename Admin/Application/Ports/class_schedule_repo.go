package ports

import (
	"context"
	entities "hadhri/Admin/Domain/Entities"
)

type IClassScheduleRepo interface {
	Create(ctx context.Context, classSchedule entities.ClassSchedule) error
	GetAll(ctx context.Context) ([]entities.ClassSchedule, error)
}
