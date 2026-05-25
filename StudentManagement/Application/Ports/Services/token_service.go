package services

type TokenService interface {
	GenerateToken(studentId int) (string, error)
}
