package usecases

import (
	"context"
	dtos "hadhri/Admin/Application/Dtos"
	ports "hadhri/Admin/Application/Ports"
)

type GetAllClassSchedules struct {
	repo ports.IClassScheduleRepo
}

func NewGetAllClassSchedulesUseCase(repo ports.IClassScheduleRepo) GetAllClassSchedules {
	return GetAllClassSchedules{repo: repo}
}

func (uc GetAllClassSchedules) Execute(ctx context.Context) ([]dtos.ClassSchedule, error) {
	rows, err := uc.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	var classSchedules []dtos.ClassSchedule

	for _, row := range rows {
		classSchedule := dtos.ClassSchedule{
			Name: row.Name,
		}

		classSchedules = append(classSchedules, classSchedule)
	}

	return classSchedules, nil
}
