package infrastructure

import (
	"context"
	ports "hadhri/Admin/Application/Ports"
	entities "hadhri/Admin/Domain/Entities"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type courseRepo struct {
	pool *pgxpool.Pool
}

func NewCourseRepo(pool *pgxpool.Pool) ports.ICourseRepo {
	return courseRepo{pool: pool}
}

func (r courseRepo) Create(ctx context.Context, course entities.Course) error {
	sql := `
		INSERT INTO courses(name)
		VALUES ($1)
	`

	_, err := r.pool.Exec(ctx, sql, course.Name)

	return err
}

func (r courseRepo) GetAll(ctx context.Context) ([]entities.Course, error) {
	sql := `
		SELECT id, name
		FROM courses
	`

	rows, err := r.pool.Query(ctx, sql)

	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[entities.Course])
}
