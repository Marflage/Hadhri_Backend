package repositories

import (
	"context"
	repositories "hadhri/StudentManagement/Application/Ports/Repositories"
	accountactivationrequest "hadhri/StudentManagement/Domain/AccountActivationRequest"

	"github.com/jackc/pgx/v5/pgxpool"
)

type accountActivationRequest struct {
	pool *pgxpool.Pool
}

func NewAccountActivationRequestRepo(pool *pgxpool.Pool) repositories.IAccountActivationRequest {
	return accountActivationRequest{pool: pool}
}

func (self accountActivationRequest) Store(ctx context.Context, e accountactivationrequest.AccountActivationRequest) (*int, error) {
	sql := `
		INSERT INTO account_activation_requests (full_name, email, phone_number, password_hash, course_plan_id, semester)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var id int

	if err := self.pool.QueryRow(ctx, sql,
		e.FullName(),
		e.Email(),
		e.PhoneNumber(),
		e.PasswordHash(),
		e.CoursePlanId(),
		e.Semester()).Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}
