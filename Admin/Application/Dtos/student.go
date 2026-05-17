package dtos

import "time"

type Student struct {
	FirstName         string    `json:"firstName" db:"first_name"`
	LastName          string    `json:"lastName" db:"last_name"`
	Email             string    `json:"email"`
	PhoneNumber       string    `json:"phoneNumber" db:"phone_number"`
	CourseName        string    `json:"courseName" db:"course_name"`
	ClassScheduleName string    `json:"classScheduleName" db:"class_schedule_name"`
	ClassSessionName  string    `json:"classSessionName" db:"class_session_name"`
	EnrollmentDate    time.Time `json:"enrollmentDate" db:"enrolled_at"`
	Semester          int       `json:"semester"`
}
