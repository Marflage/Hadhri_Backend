package ports

import (
	"context"
	entities "hadhri/Admin/Domain/Entities"
)

type IClassSessionRepo interface {
	Create(ctx context.Context, classSession entities.ClassSession) error
	GetAll(ctx context.Context) ([]entities.ClassSession, error)
}
