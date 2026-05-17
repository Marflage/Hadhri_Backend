package queryservices

import (
	"context"
	dtos "hadhri/Admin/Application/Dtos"
	queryservices "hadhri/Admin/Application/Ports/QueryServices"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type student struct {
	pool *pgxpool.Pool
}

func NewStudentQueryService(pool *pgxpool.Pool) queryservices.IStudent {
	return student{pool: pool}
}

func (qs student) Get(ctx context.Context, id int) (dtos.Student, error) {
	sql := `
		SELECT s.first_name,
			s.last_name,
			s.email,
			s.phone_number,
			e.enrolled_at,
			e.semester,
			c.name AS course_name,
			cs.name AS class_schedule_name,
			cses.name AS class_session_name
		FROM students s
			INNER JOIN public.enrollments e ON s.id = e.student_id
			INNER JOIN public.course_plans cp ON cp.id = e.course_plan_id
			INNER JOIN public.courses c ON c.id = cp.course_id
			INNER JOIN public.class_schedules cs ON cs.id = cp.class_schedule_id
			INNER JOIN public.class_sessions cses ON cses.id = cp.class_session_id
		WHERE s.id = $1
	`

	rows, err := qs.pool.Query(ctx, sql, id)

	if err != nil {
		// TODO: Should fmt.Errorf() be used?
		return dtos.Student{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[dtos.Student])
}
