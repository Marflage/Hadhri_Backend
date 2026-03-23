package responses

type GetStudentDetails struct {
	StudentId         int    `json:"studentId"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	CourseName        string `json:"courseName"`
	ClassScheduleName string `json:"classScheduleName"`
	ClassSessionName  string `json:"classSessionName"`
	Semester          int    `json:"semester"`
}
