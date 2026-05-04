package usecases

import (
	"context"
	ports "hadhri/Admin/Application/Ports"
	querymodels "hadhri/Admin/Domain/QueryModels"
)

type GetAllCoursePlans struct {
	repo ports.ICoursePlanRepo
}

func NewGetAllCoursePlansUseCase(repo ports.ICoursePlanRepo) GetAllCoursePlans {
	return GetAllCoursePlans{repo: repo}
}

func (uc GetAllCoursePlans) Execute(ctx context.Context) ([]querymodels.CoursePlan, error) {
	rows, err := uc.repo.GetAll(ctx)

	return rows, err
}
