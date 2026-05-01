package usecases

import (
	"context"
	ports "hadhri/Admin/Application/Ports"
	dtos "hadhri/Admin/Domain/Dtos"
)

type GetAllClassSessions struct {
	repo ports.IClassSessionRepo
}

func NewGetAllClassSessionsUseCase(repo ports.IClassSessionRepo) GetAllClassSessions {
	return GetAllClassSessions{repo: repo}
}

func (uc GetAllClassSessions) Execute(ctx context.Context) ([]dtos.ClassSession, error) {
	rows, err := uc.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	var classSessions []dtos.ClassSession

	for _, row := range rows {
		classSession := dtos.ClassSession{
			Name: row.Name,
		}

		classSessions = append(classSessions, classSession)
	}

	return classSessions, nil
}
