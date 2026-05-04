package infrastructure

import (
	"context"
	ports "hadhri/Admin/Application/Ports"
	querymodels "hadhri/Admin/Domain/QueryModels"
	dbmodels "hadhri/Admin/Infrastructure/DbModels"

	"github.com/jackc/pgx/v5"
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

func (r coursePlanRepo) GetAll(ctx context.Context) ([]querymodels.CoursePlan, error) {
	sql := `
	SELECT cp.id AS Id, c.name AS CourseName, cs.name AS ClassScheduleName, s.name AS ClassSessionName, cp.is_active AS IsActive
	FROM course_plans cp
			INNER JOIN public.courses c ON c.id = cp.course_id
			INNER JOIN public.class_schedules cs ON cs.id = cp.class_schedule_id
			INNER JOIN public.class_sessions s ON s.id = cp.class_session_id
	`

	rows, err := r.pool.Query(ctx, sql)

	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[querymodels.CoursePlan])
}
