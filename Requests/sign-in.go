package requests

type SignInRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
