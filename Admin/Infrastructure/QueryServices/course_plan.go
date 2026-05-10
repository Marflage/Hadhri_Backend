package queryservices

import (
	"context"
	dtos "hadhri/Admin/Application/Dtos"
	queryservices "hadhri/Admin/Application/Ports/QueryServices"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type coursePlan struct {
	pool *pgxpool.Pool
}

func NewCoursePlanQueryService(pool *pgxpool.Pool) queryservices.ICoursePlan {
	return coursePlan{pool: pool}
}

func (qs coursePlan) GetAll(ctx context.Context) ([]dtos.CoursePlan, error) {
	sql := `
		SELECT cp.id        AS Id,
		c.id         AS CourseId,
		c.name       AS CourseName,
		cs.id        AS ClassScheduleId,
		cs.name      AS ClassScheduleName,
		s.id         AS ClassSessionId,
		s.name       AS ClassSessionName,
		cp.is_active AS IsActive
		FROM course_plans cp
				INNER JOIN public.courses c ON c.id = cp.course_id
				INNER JOIN public.class_schedules cs ON cs.id = cp.class_schedule_id
				INNER JOIN public.class_sessions s ON s.id = cp.class_session_id
	`

	rows, err := qs.pool.Query(ctx, sql)

	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[dtos.CoursePlan])
}
