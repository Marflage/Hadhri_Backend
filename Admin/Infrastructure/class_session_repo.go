package infrastructure

import (
	"context"
	ports "hadhri/Admin/Application/Ports"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type classSessionRepo struct {
	pool *pgxpool.Pool
}

func NewClassSessionRepo(pool *pgxpool.Pool) ports.IClassSessionRepo {
	return classSessionRepo{pool: pool}
}

func (r classSessionRepo) Create(ctx context.Context, entity dbmodels.ClassSession) error {
	cmd := `
		INSERT INTO class_sessions(name, start_time, end_time)
		VALUES ($1, $2, $3);
	`

	_, err := r.pool.Exec(ctx, cmd, entity.Name, entity.StartTime, entity.EndTime)

	return err
}

func (r classSessionRepo) GetAll(ctx context.Context) ([]dbmodels.ClassSession, error) {
	sql := `
		SELECT id, name, start_time, end_time
		FROM class_sessions
	`

	rows, err := r.pool.Query(ctx, sql)

	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[dbmodels.ClassSession])
}
