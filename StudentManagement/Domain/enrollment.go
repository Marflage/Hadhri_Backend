package domain

import (
	"errors"
	"time"
)

type Enrollment struct {
	CoursePlanId int
	Semester     int
	enrolledAt   time.Time
}

func NewEnrollment(coursePlanId int, semester int) (Enrollment, error) {
	if coursePlanId < 1 {
		// TODO: Why is an empty instance returned instead of nil?
		return Enrollment{}, errors.New("Invalid course plan ID.")
	}

	if semester < 1 {
		return Enrollment{}, errors.New("Invalid semester.")
	}

	return Enrollment{
		CoursePlanId: coursePlanId,
		Semester:     semester,
		enrolledAt:   time.Now(),
	}, nil
}
