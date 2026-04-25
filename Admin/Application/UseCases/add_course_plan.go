package usecases

import (
	"context"
	commands "hadhri/Admin/Application/Commands"
	ports "hadhri/Admin/Application/Ports"
	domain "hadhri/Admin/Domain"
)

type AddCoursePlan struct {
	repo ports.ICoursePlanRepo
}

func NewAddCoursePlanUseCase(repo ports.ICoursePlanRepo) AddCoursePlan {
	return AddCoursePlan{repo: repo}
}

func (uc AddCoursePlan) Execute(ctx context.Context, cmd commands.AddCoursePlan) error {
	// TODO: Validate for business rules.
	// TODO: Check if course, class schedule, and class session with the passed IDs exist in the database.

	entity := domain.CoursePlan{
		CourseId:        cmd.CourseId,
		ClassScheduleId: cmd.ClassScheduleId,
		ClassSessionId:  cmd.ClassSessionId,
		IsActive:        cmd.IsActive,
	}

	return uc.repo.Create(ctx, entity)
}
