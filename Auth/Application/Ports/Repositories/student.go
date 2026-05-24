package repositories

import (
	"context"
	student "hadhri/Auth/Domain/Student"
)

type IStudent interface {
	SignUp(ctx context.Context, student student.Student) (*int, error)
}
