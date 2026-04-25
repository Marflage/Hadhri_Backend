package domain

import "time"

type CoursePlan struct {
	id              int
	insertedAt      time.Time
	updatedAt       time.Time
	CourseId        int
	ClassScheduleId int
	ClassSessionId  int
	IsActive        bool
}
