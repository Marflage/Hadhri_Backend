package entities

import "time"

type CoursePlan struct {
	Id              int
	insertedAt      time.Time
	updatedAt       time.Time
	CourseId        int
	ClassScheduleId int
	ClassSessionId  int
	IsActive        bool
}
