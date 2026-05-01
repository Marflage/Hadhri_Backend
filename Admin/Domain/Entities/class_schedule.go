package entities

import "time"

type ClassSchedule struct {
	id         int
	insertedAt time.Time
	updatedAt  time.Time
	Name       string
}
