package dtos

type CoursePlan struct {
	Id                int
	CourseId          int
	CourseName        string
	ClassScheduleId   int
	ClassScheduleName string
	ClassSessionId    int
	ClassSessionName  string
	IsActive          bool
}
