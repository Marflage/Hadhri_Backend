package responses

import "time"

type GetStudentEnrollmentDetails struct {
	CourseName            string    `json:"courseName"`
	ClassScheduleName     string    `json:"classScheduleName"`
	ClassSessionName      string    `json:"classSessionName"`
	ClassSessionStartTime time.Time `json:"classSessionStartTime"`
	ClassSessionEndTime   time.Time `json:"classSessionEndTime"`
	Semester              int       `json:"semester"`
}
