package ports

import (
	"context"
	domain "hadhri/Admin/Domain"
)

type IClassScheduleRepo interface {
	Create(ctx context.Context, classSchedule domain.ClassSchedule) error
}
