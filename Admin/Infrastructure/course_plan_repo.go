package infrastructure

import (
	"context"
	ports "hadhri/Admin/Application/Ports"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"

	"github.com/jackc/pgx/v5/pgxpool"
)

type coursePlanRepo struct {
	pool *pgxpool.Pool
}

func NewCoursePlanRepo(pool *pgxpool.Pool) ports.ICoursePlanRepo {
	return coursePlanRepo{pool: pool}
}

func (r coursePlanRepo) Create(ctx context.Context, entity dbmodels.CoursePlan) error {
	sql := `
		INSERT INTO course_plans(course_id, class_schedule_id, class_session_id, is_active)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.pool.Exec(ctx, sql, entity.CourseId, entity.ClassScheduleId, entity.ClassSessionId, entity.IsActive)

	return err
}
