package querymodels

type GetStudentDetails struct {
	StudentId         int
	FirstName         string
	LastName          string
	CourseName        string
	ClassScheduleName string
	ClassSessionName  string
	Semester          int
}
