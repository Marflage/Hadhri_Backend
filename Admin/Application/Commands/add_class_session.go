package commands

import "time"

type AddClassSession struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time
}
