package dtos

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	StudentId int `json:"student_id"`
	jwt.RegisteredClaims
}
