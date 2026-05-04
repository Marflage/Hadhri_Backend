package dbmodels

import "time"

type ClassSchedule struct {
	Id         int
	insertedAt time.Time
	updatedAt  time.Time
	Name       string
}
