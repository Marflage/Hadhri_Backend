package usecases

import (
	"context"
	dtos "hadhri/Admin/Application/Dtos"
	queryservices "hadhri/Admin/Application/Ports/QueryServices"
)

type GetStudent struct {
	queryService queryservices.IStudent
}

func NewGetStudentUseCase(queryService queryservices.IStudent) GetStudent {
	return GetStudent{queryService: queryService}
}

func (uc GetStudent) Execute(ctx context.Context, id int) (dtos.Student, error) {
	dto, err := uc.queryService.Get(ctx, id)

	return dto, err
}
