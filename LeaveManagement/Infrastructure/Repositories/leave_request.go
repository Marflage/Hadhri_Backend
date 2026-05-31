package repositories

import (
	"context"
	repositories "hadhri/LeaveManagement/Application/Ports/Repositories"
	domain "hadhri/LeaveManagement/Domain/LeaveRequest"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type leaveRequest struct {
	pool *pgxpool.Pool
}

func NewLeaveRequestRepo(pool *pgxpool.Pool) repositories.ILeaveRequest {
	return leaveRequest{pool: pool}
}

func (self leaveRequest) AddLeaveRequest(ctx context.Context, entity domain.LeaveRequest) error {
	sql := `
		INSERT INTO leave_requests(student_id, start_date, end_date, reason)
		VALUES ($1, $2, $3, $4);
	`

	_, err := self.pool.Exec(ctx, sql,
		entity.GetStudentId(),
		entity.GetStartDate(),
		entity.GetEndDate(),
		entity.GetReason())

	return err
}

func (self leaveRequest) AlreadyExists(ctx context.Context, e domain.LeaveRequest) error {
	sql := `
		SELECT EXISTS(SELECT 1
              FROM leave_requests
              WHERE student_id = $1
			  	AND start_date <= $2
                AND end_date >= $3)
	`

	exists := false

	return self.pool.QueryRow(ctx, sql,
		e.GetStudentId(),
		e.GetStartDate(),
		e.GetEndDate()).Scan(&exists)
}

func (self leaveRequest) Update(ctx context.Context, e domain.LeaveRequest) error {
	sql := `
		UPDATE leave_requests
		SET start_date = $1,
		end_date   = $2,
		reason     = $3
		WHERE id = $4
	`

	_, err := self.pool.Exec(ctx, sql, e.GetStartDate(), e.GetEndDate(), e.GetReason(), e.GetId())

	return err
}

func (self leaveRequest) Get(ctx context.Context, id uint) (*domain.LeaveRequest, error) {
	sql := `
		SELECT id, student_id, start_date, end_date, reason, status
		FROM leave_requests
		WHERE id = $1
	`

	type dbRow struct {
		Id        int       `db:"id"`
		StudentId int       `db:"student_id"`
		StartDate time.Time `db:"start_date"`
		EndDate   time.Time `db:"end_date"`
		Reason    string    `db:"reason"`
		Status    string    `db:"status"`
	}

	rows, err := self.pool.Query(ctx, sql, id)

	if err != nil {
		return nil, err
	}

	row, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbRow])

	if err != nil {
		return nil, err
	}

	e := domain.ReconstituteLeaveRequest(uint(row.Id), uint(row.StudentId), row.StartDate, row.EndDate, row.Reason, row.Status)

	return &e, nil
}

func (self leaveRequest) Cancel(ctx context.Context, id uint, studentId uint) error {
	sql := `
		UPDATE leave_requests
		SET status = 'canceled',
		status_changed_at = CURRENT_TIMESTAMP
		WHERE id = $1
		AND student_id = $2
	`

	_, err := self.pool.Exec(ctx, sql, id, studentId)

	return err
}

func (self leaveRequest) Exists(ctx context.Context, id uint, studentId uint) (*bool, error) {
	sql := `
		SELECT EXISTS(SELECT 1
				FROM leave_requests
				WHERE id = $1
					AND student_id = $2)
	`

	exists := false

	if err := self.pool.QueryRow(ctx, sql, id, studentId).Scan(&exists); err != nil {
		return nil, err
	}

	return &exists, nil
}
