package domain

import "time"

type Course struct {
	id         int
	insertedAt time.Time
	updatedAt  time.Time
	Name       string
}
