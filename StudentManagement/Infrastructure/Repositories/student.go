package repositories

import (
	"context"
	repositories "hadhri/StudentManagement/Application/Ports/Repositories"
	student "hadhri/StudentManagement/Domain/Student"

	"github.com/jackc/pgx/v5/pgxpool"
)

type studentRepo struct {
	pool *pgxpool.Pool
}

func NewStudentRepo(pool *pgxpool.Pool) repositories.IStudent {
	return studentRepo{pool: pool}
}

// TODO: A CTE can be used for atomicity and performance. Research.
func (r studentRepo) SignUp(ctx context.Context, student student.Student) (*int, error) {
	tx, err := r.pool.Begin(ctx)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	sql := `
		INSERT INTO students(full_name, email, phone_number, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var studentId int

	if err := tx.QueryRow(ctx, sql,
		student.GetFullName(),
		student.GetEmail(),
		student.GetPhoneNumber(),
		student.GetPassword(),
	).Scan(&studentId); err != nil {
		return nil, err
	}

	sql2 := `
		INSERT INTO enrollments(student_id, course_plan_id, semester)
		VALUES ($1, $2, $3)
	`

	if _, err := tx.Exec(ctx, sql2,
		studentId,
		student.GetCoursePlanId(),
		student.GetSemester()); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &studentId, nil
}
