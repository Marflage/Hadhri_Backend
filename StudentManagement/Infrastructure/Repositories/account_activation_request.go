package repositories

import (
	"context"
	repositories "hadhri/StudentManagement/Application/Ports/Repositories"
	accountactivationrequest "hadhri/StudentManagement/Domain/AccountActivationRequest"
	dbmodels "hadhri/StudentManagement/Infrastructure/DbModels"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type accountActivationRequest struct {
	pool *pgxpool.Pool
}

func NewAccountActivationRequestRepo(pool *pgxpool.Pool) repositories.IAccountActivationRequest {
	return accountActivationRequest{pool: pool}
}

func (self accountActivationRequest) Save(ctx context.Context, e accountactivationrequest.AccountActivationRequest) (*int, error) {
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

func (self accountActivationRequest) Exists(ctx context.Context, id uint) (*bool, error) {
	sql := `
		SELECT EXISTS(SELECT 1
				FROM account_activation_requests
				WHERE id = $1
					AND status = 'pending')
	`

	exists := false

	if err := self.pool.QueryRow(ctx, sql, id).Scan(&exists); err != nil {
		return nil, err
	}

	return &exists, nil
}

func (self accountActivationRequest) Approve(ctx context.Context, id uint) error {
	sql := `
		UPDATE account_activation_requests
		SET status            = 'approved',
			status_changed_at = CURRENT_TIMESTAMP
		WHERE id = $1
		AND status = 'pending'
	`

	if _, err := self.pool.Exec(ctx, sql, id); err != nil {
		return err
	}

	row, err := self.get(ctx, id)

	if err != nil {
		return err
	}

	if err := self.signUp(ctx, *row); err != nil {
		return err
	}

	return nil
}

func (self accountActivationRequest) Decline(ctx context.Context, id uint) error {
	sql := `
		UPDATE account_activation_requests
		SET status            = 'declined',
			status_changed_at = CURRENT_TIMESTAMP
		WHERE id = $1
		AND status = 'pending'
	`

	if _, err := self.pool.Exec(ctx, sql, id); err != nil {
		return err
	}

	return nil
}

func (self accountActivationRequest) get(ctx context.Context, id uint) (*dbmodels.AccountActivationRequest, error) {
	sql := `
		SELECT id,
		inserted_at,
		full_name,
		email,
		phone_number,
		password_hash,
		course_plan_id,
		semester,
		status_changed_at
		FROM account_activation_requests
		WHERE id = $1
	`

	rows, err := self.pool.Query(ctx, sql, id)

	if err != nil {
		return nil, err
	}

	row, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[dbmodels.AccountActivationRequest])

	if err != nil {
		return nil, err
	}

	return &row, nil
}

// TODO: A CTE can be used for atomicity and performance. Research.
func (self accountActivationRequest) signUp(ctx context.Context, row dbmodels.AccountActivationRequest) error {
	tx, err := self.pool.Begin(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	sql := `
		INSERT INTO students(id, full_name, email, phone_number, password_hash)
		VALUES ($1, $2, $3, $4, $5)
	`

	if _, err := tx.Exec(ctx, sql,
		row.Id,
		row.FullName,
		row.Email,
		row.PhoneNumber,
		row.PasswordHash); err != nil {
		return err
	}

	sql2 := `
		INSERT INTO enrollments(student_id, course_plan_id, enrolled_at, semester)
		VALUES ($1, $2, CURRENT_TIMESTAMP, $3)
	`

	if _, err := tx.Exec(ctx, sql2,
		row.Id,
		row.CoursePlanId,
		row.Semester); err != nil {
		return err
	}

	sql3 := `
		INSERT INTO student_account_activation_approvals(student_id, applied_at, approved_at)
		VALUES ($1, $2, $3);
	`

	if _, err := tx.Exec(ctx, sql3,
		row.Id,
		row.InsertedAt,
		row.StatusChangedAt); err != nil {
		return err
	}

	// TODO: Delete the activation request from account_activation_requests table.

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
