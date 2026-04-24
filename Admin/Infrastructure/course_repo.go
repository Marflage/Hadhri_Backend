package infrastructure

import (
	"context"
	ports "hadhri/Admin/Application/Ports"
	domain "hadhri/Admin/Domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type courseRepo struct {
	pool *pgxpool.Pool
}

func NewCourseRepo(pool *pgxpool.Pool) ports.ICourseRepo {
	return courseRepo{pool: pool}
}

// TODO: is it necessary for the receiver to be a pointer?
func (r courseRepo) Create(ctx context.Context, course domain.Course) error {
	cmd := `
		INSERT INTO courses(name)
		VALUES ($1)
	`

	_, err := r.pool.Exec(ctx, cmd, course.Name)

	return err
}
