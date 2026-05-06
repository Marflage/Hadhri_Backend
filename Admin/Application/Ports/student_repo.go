package ports

import (
	"context"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"
)

type IStudentRepo interface {
	Add(ctx context.Context, row dbmodels.Student) error
}
