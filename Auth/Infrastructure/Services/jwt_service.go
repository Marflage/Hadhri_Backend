package services

import (
	services "hadhri/Auth/Application/Ports/Services"
	constants "hadhri/Auth/Infrastructure/Constants"
	dtos "hadhri/Auth/Infrastructure/Dtos"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct{}

func NewJwtService() services.TokenService {
	return jwtService{}
}

func (s jwtService) GenerateToken(studentId int) (string, error) {
	claims := dtos.CustomClaims{
		StudentId: studentId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    constants.Issuer,
			Audience:  jwt.ClaimStrings{constants.Audience},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(constants.ExpirationMinutes) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(constants.Jwtkey)
}
