package commands

import "time"

type RequestLeave struct {
	StudentId uint
	StartDate time.Time
	EndDate   time.Time
	Reason    string
}
