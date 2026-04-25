package infrastructure

import (
	"context"
	ports "hadhri/Admin/Application/Ports"
	domain "hadhri/Admin/Domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type classSessionRepo struct {
	pool *pgxpool.Pool
}

func NewClassSessionRepo(pool *pgxpool.Pool) ports.IClassSessionRepo {
	return &classSessionRepo{pool: pool}
}

func (r *classSessionRepo) Create(ctx context.Context, entity domain.ClassSession) error {
	cmd := `
		INSERT INTO class_sessions(name, start_time, end_time)
		VALUES ($1, $2, $3);
	`

	_, err := r.pool.Exec(ctx, cmd, entity.Name, entity.StartTime, entity.EndTime)

	return err
}
