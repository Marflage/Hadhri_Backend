package repositories

import (
	"context"
	student "hadhri/StudentManagement/Domain/Student"
)

type IStudent interface {
	SignUp(ctx context.Context, student student.Student) (*uint, error)
}
