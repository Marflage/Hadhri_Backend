package usecases

import (
	"context"
	commands "hadhri/Admin/Application/Commands"
	ports "hadhri/Admin/Application/Ports"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"
)

type AddClassSchedule struct {
	repo ports.IClassScheduleRepo
}

func NewAddClassScheduleUseCase(repo ports.IClassScheduleRepo) AddClassSchedule {
	return AddClassSchedule{repo: repo}
}

func (uc AddClassSchedule) Execute(ctx context.Context, cmd commands.AddClassSchedule) error {
	classSchedule := dbmodels.ClassSchedule{
		Name: cmd.Name,
	}

	return uc.repo.Create(ctx, classSchedule)
}
