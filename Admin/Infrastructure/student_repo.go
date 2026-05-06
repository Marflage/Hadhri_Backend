package infrastructure

import (
	"context"
	ports "hadhri/Admin/Application/Ports"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"

	"github.com/jackc/pgx/v5/pgxpool"
)

type studentRepo struct {
	pool *pgxpool.Pool
}

func NewStudentRepo(pool *pgxpool.Pool) ports.IStudentRepo {
	return studentRepo{pool: pool}
}

func (r studentRepo) Add(ctx context.Context, row dbmodels.Student) error {
	sql := `
		INSERT INTO students (first_name, last_name, email, phone_number, password)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.pool.Exec(ctx, sql, row.FirstName, row.LastName, row.Email, row.PhoneNumber, row.Password)

	return err
}
