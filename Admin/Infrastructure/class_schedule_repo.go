package infrastructure

import (
	"context"
	ports "hadhri/Admin/Application/Ports"
	domain "hadhri/Admin/Domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type classScheduleRepo struct {
	pool *pgxpool.Pool
}

func NewClassScheduleRepo(pool *pgxpool.Pool) ports.IClassScheduleRepo {
	return &classScheduleRepo{pool: pool}
}

func (r *classScheduleRepo) Create(ctx context.Context, classSchedule domain.ClassSchedule) error {
	cmd := `
		INSERT INTO class_schedules(name)
		VALUES ($1)
	`

	_, err := r.pool.Exec(ctx, cmd, classSchedule.Name)

	return err
}
