package commands

import "time"

type EditLeaveRequest struct {
	Id        uint
	StartDate *time.Time
	EndDate   *time.Time
	Reason    *string
}
