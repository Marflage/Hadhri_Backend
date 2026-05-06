package dbmodels

import "time"

type Student struct {
	Id          int
	insertedAt  time.Time
	updatedAt   time.Time
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Password    string
}
