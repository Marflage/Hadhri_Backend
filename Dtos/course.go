package dtos

type Course struct {
	Id             int             `json:"id"`
	Name           string          `json:"name"`
	ClassSchedules []ClassSchedule `json:"classSchedules"`
}
