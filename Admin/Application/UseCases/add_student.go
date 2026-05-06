package usecases

import (
	"context"
	commands "hadhri/Admin/Application/Commands"
	ports "hadhri/Admin/Application/Ports"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"
)

type AddStudent struct {
	repo ports.IStudentRepo
}

func NewAddStudentUseCase(repo ports.IStudentRepo) AddStudent {
	return AddStudent{repo: repo}
}

func (uc AddStudent) Execute(ctx context.Context, cmd commands.AddStudent) error {
	// TODO: Validate for business rules.

	// TODO: Add enrollment as well.

	row := dbmodels.Student{
		FirstName:   cmd.FirstName,
		LastName:    cmd.LastName,
		Email:       cmd.Email,
		PhoneNumber: cmd.PhoneNumber,
		Password:    cmd.Password,
	}

	return uc.repo.Add(ctx, row)
}
