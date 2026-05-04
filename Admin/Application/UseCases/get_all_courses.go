package usecases

import (
	"context"
	dtos "hadhri/Admin/Application/Dtos"
	ports "hadhri/Admin/Application/Ports"
)

type GetAllCourses struct {
	repo ports.ICourseRepo
}

func NewGetAllCoursesUseCase(repo ports.ICourseRepo) GetAllCourses {
	return GetAllCourses{repo: repo}
}

func (uc GetAllCourses) Execute(ctx context.Context) ([]dtos.Course, error) {
	rows, err := uc.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	var courses []dtos.Course

	for _, row := range rows {
		course := dtos.Course{
			Name: row.Name,
		}

		courses = append(courses, course)
	}

	return courses, nil
}
