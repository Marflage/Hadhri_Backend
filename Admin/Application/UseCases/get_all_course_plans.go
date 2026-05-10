package usecases

import (
	"context"
	dtos "hadhri/Admin/Application/Dtos"
	queryservices "hadhri/Admin/Application/Ports/QueryServices"
)

type GetAllCoursePlans struct {
	queryService queryservices.ICoursePlan
}

func NewGetAllCoursePlansUseCase(queryService queryservices.ICoursePlan) GetAllCoursePlans {
	return GetAllCoursePlans{queryService: queryService}
}

func (uc GetAllCoursePlans) Execute(ctx context.Context) ([]dtos.CoursePlan, error) {
	rows, err := uc.queryService.GetAll(ctx)

	return rows, err
}
