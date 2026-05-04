package infrastructure

import (
	"context"
	ports "hadhri/Admin/Application/Ports"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type classScheduleRepo struct {
	pool *pgxpool.Pool
}

func NewClassScheduleRepo(pool *pgxpool.Pool) ports.IClassScheduleRepo {
	return classScheduleRepo{pool: pool}
}

func (r classScheduleRepo) Create(ctx context.Context, classSchedule dbmodels.ClassSchedule) error {
	sql := `
		INSERT INTO class_schedules(name)
		VALUES ($1)
	`

	_, err := r.pool.Exec(ctx, sql, classSchedule.Name)

	return err
}

func (r classScheduleRepo) GetAll(ctx context.Context) ([]dbmodels.ClassSchedule, error) {
	sql := `
		SELECT id, name
		FROM class_schedules
	`

	rows, err := r.pool.Query(ctx, sql)

	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[dbmodels.ClassSchedule])
}
