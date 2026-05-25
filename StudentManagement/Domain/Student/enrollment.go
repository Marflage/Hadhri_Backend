package student

import (
	"fmt"
	"time"
)

type enrollment struct {
	coursePlanId int
	semester     int
	enrolledAt   time.Time
}

func newEnrollment(coursePlanId int, semester int) (enrollment, error) {
	if coursePlanId < 1 {
		return enrollment{}, fmt.Errorf("Invalid course plan ID.")
	}

	if semester < 1 {
		return enrollment{}, fmt.Errorf("Invalid semester.")
	}

	return enrollment{
		coursePlanId: coursePlanId,
		semester:     semester,
		enrolledAt:   time.Now(),
	}, nil
}
