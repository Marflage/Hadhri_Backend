package responses

type SignIn struct {
	StudentId int    `json:"studentId"`
	Token     string `json:"token"`
}
