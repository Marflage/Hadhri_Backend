type ClassSchedule struct {
	Id            int            `json:"id"`
	Name          string         `json:"name"`
	ClassSessions []ClassSession `json:"classSessions"`
}