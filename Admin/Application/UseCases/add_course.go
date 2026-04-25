package usecases

import (
	"context"
	commands "hadhri/Admin/Application/Commands"
	ports "hadhri/Admin/Application/Ports"
	domain "hadhri/Admin/Domain"
)

// TODO: Should this use case have an interface so that it can be unexported to enforce constructor invocation?
type AddCourse struct {
	repo ports.ICourseRepo
}

// TODo: Is it necessary to return a pointer?
func NewAddCourseUseCase(repo ports.ICourseRepo) AddCourse {
	return AddCourse{repo: repo}
}

// TODO: Should the receiver be a pointer?
// TODO: Why should error be returned from this method?
func (uc AddCourse) Execute(ctx context.Context, cmd commands.AddCourse) error {
	// TODO: Validate input/command as this method is not necessarily called only from the web API layer.

	// TODO: Check for any duplicate or exisiting course with the same name.

	course := domain.Course{Name: cmd.Name}

	if err := uc.repo.Create(ctx, course); err != nil {
		return err
	}

	return nil
}
