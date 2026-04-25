package ports

import (
	"context"
	domain "hadhri/Admin/Domain"
)

type IClassSessionRepo interface {
	Create(ctx context.Context, classSession domain.ClassSession) error
}
