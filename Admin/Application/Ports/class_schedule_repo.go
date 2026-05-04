package ports

import (
	"context"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"
)

type IClassScheduleRepo interface {
	Create(ctx context.Context, classSchedule dbmodels.ClassSchedule) error
	GetAll(ctx context.Context) ([]dbmodels.ClassSchedule, error)
}
