package querymodels

type GetCoursePlans struct {
	CourseId          int
	CourseName        string
	ClassScheduleId   int
	ClassScheduleName string
	ClassSessionId    int
	ClassSessionName  string
	AvailableSemester int
}
