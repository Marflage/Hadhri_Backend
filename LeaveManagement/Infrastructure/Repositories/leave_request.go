package repositories

import (
	"context"
	repositories "hadhri/LeaveManagement/Application/Ports/Repositories"
	domain "hadhri/LeaveManagement/Domain/LeaveRequest"

	"github.com/jackc/pgx/v5/pgxpool"
)

type leaveRequest struct {
	pool *pgxpool.Pool
}

func NewLeaveRequestRepo(pool *pgxpool.Pool) repositories.ILeaveRequest {
	return leaveRequest{pool: pool}
}

func (r leaveRequest) AddLeaveRequest(ctx context.Context, entity domain.LeaveRequest) error {
	sql := `
		INSERT INTO leave_requests(student_id, start_date, end_date, reason)
		VALUES ($1, $2, $3, $4);
	`

	_, err := r.pool.Exec(ctx, sql,
		entity.GetStudentId(),
		entity.GetStartDate(),
		entity.GetEndDate(),
		entity.GetReason())

	return err
}

func (r leaveRequest) AlreadyExists(ctx context.Context, entity domain.LeaveRequest) error {
	sql := `
		SELECT EXISTS(SELECT 1
              FROM leave_requests
              WHERE student_id = $1
                AND start_date <= $2
                AND end_date >= $3)
	`

	exists := false

	return r.pool.QueryRow(ctx, sql,
		entity.GetStudentId(),
		entity.GetStartDate(),
		entity.GetEndDate()).Scan(&exists)
}
