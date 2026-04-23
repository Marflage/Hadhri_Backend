package responses

type GetStudentDetails struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
