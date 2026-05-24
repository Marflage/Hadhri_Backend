package queryservices

import (
	"context"
	queryservices "hadhri/Auth/Application/Ports/QueryServices"

	"github.com/jackc/pgx/v5/pgxpool"
)

type coursePlan struct {
	pool *pgxpool.Pool
}

func NewCoursePlanQueryService(pool *pgxpool.Pool) queryservices.ICoursePlan {
	return coursePlan{pool: pool}
}

func (r coursePlan) GetId(ctx context.Context, courseId int, classScheduleId int, classSessionId int, isActive bool, semester int) (int, error) {
	sql := `
		SELECT cp.id
		FROM course_plans cp
			JOIN available_semesters a
				ON cp.id = a.course_plan_id
		WHERE course_id = $1
		AND cp.class_schedule_id = $2
		AND cp.class_session_id = $3
		AND cp.is_active = TRUE
		AND a.semester = $4
		`

	var coursePlanId int

	err := r.pool.QueryRow(
		ctx,
		sql,
		courseId,
		classScheduleId,
		classSessionId,
		semester,
	).Scan(&coursePlanId)

	return coursePlanId, err
}
