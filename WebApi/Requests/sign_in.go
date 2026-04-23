package requests

type SignIn struct {
	Email    string `json:"email" binding:"required,noBlank,email"`
	Password string `json:"password" binding:"required,noBlank"`
}
