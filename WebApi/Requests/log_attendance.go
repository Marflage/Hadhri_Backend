package requests

type LogAttendance struct {
	StudentId int `form:"studentId" binding:"gte=1"`
}
