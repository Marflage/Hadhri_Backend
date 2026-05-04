package usecases

import (
	"context"
	"errors"
	commands "hadhri/Admin/Application/Commands"
	ports "hadhri/Admin/Application/Ports"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"
	"time"
)

type AddClassSession struct {
	repo ports.IClassSessionRepo
}

func NewAddClassSessionUseCase(repo ports.IClassSessionRepo) AddClassSession {
	// TODO: How to enforce setting repo in the initializer?
	return AddClassSession{repo: repo}
}

func (uc AddClassSession) Execute(ctx context.Context, cmd commands.AddClassSession) error {
	// TODO: Validate incoming data

	cmd.StartTime = time.Date(0000, 01, 01, cmd.StartTime.Hour(), cmd.StartTime.Minute(), 0, 0, time.Local)
	cmd.EndTime = time.Date(0000, 01, 01, cmd.EndTime.Hour(), cmd.EndTime.Minute(), 0, 0, time.Local)

	if !cmd.EndTime.After(cmd.StartTime) {
		return errors.New("End time must be after start time.")
	}

	entity := dbmodels.ClassSession{
		Name:      cmd.Name,
		StartTime: cmd.StartTime,
		EndTime:   cmd.EndTime,
	}

	err := uc.repo.Create(ctx, entity)

	return err
}
