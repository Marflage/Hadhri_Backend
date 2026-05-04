package dbmodels

import "time"

type Course struct {
	Id         int
	insertedAt time.Time
	updatedAt  time.Time
	Name       string
}
