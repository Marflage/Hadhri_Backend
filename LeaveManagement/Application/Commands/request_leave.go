package commands

import "time"

type RequestLeave struct {
	StudentId int
	StartDate time.Time
	EndDate   time.Time
	Reason    string
}
