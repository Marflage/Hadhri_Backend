package responses

type Auth struct {
	StudentId int    `json:"studentId"`
	Token     string `json:"token"`
}
