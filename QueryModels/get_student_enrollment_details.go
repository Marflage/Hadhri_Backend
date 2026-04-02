package querymodels

import "time"

type GetStudentEnrollmentDetails struct {
	CourseName            string
	ClassScheduleName     string
	ClassSessionName      string
	ClassSessionStartTime time.Time
	ClassSessionEndTime   time.Time
	Semester              int
}
