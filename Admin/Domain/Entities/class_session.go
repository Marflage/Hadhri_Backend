package entities

import "time"

type ClassSession struct {
	id         int
	insertedAt time.Time
	updatedAt  time.Time
	Name       string
	StartTime  time.Time
	EndTime    time.Time
}
