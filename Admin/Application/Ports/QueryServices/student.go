package queryservices

import (
	"context"
	dtos "hadhri/Admin/Application/Dtos"
)

type IStudent interface {
	Get(ctx context.Context, id int) (dtos.Student, error)
}
