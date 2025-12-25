package entities

import "time"

type CoursePlan struct {
	Id              int
	InsertedAt      time.Time
	UpdatedAt       time.Time
	CourseId        int
	ClassScheduleId int
	ClassSessionId  int
	IsActive        bool
}
