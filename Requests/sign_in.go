package requests

type SignInRequest struct {
	Email    string `json:"email" binding:"required,noBlank,email"`
	Password string `json:"password" binding:"required,noBlank"`
}
