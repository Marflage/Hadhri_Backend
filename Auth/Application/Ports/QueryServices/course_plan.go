package queryservices

import "context"

type ICoursePlan interface {
	GetId(ctx context.Context, courseId int, classScheduleId int, classSessionId int, isActive bool, semester int) (int, error)
}
