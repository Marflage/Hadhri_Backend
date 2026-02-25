package requests

type LogAttendance struct {
	StudentId int `json:"studentId" binding:"required,gte=1"`
}
