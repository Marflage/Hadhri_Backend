package repositories

import (
	"context"
	domain "hadhri/LeaveManagement/Domain/LeaveRequest"
)

type ILeaveRequest interface {
	AddLeaveRequest(ctx context.Context, entity domain.LeaveRequest) error
	AlreadyExists(ctx context.Context, entity domain.LeaveRequest) error
	Update(ctx context.Context, entity domain.LeaveRequest) error
	Get(ctx context.Context, id uint, studentId uint) (*domain.LeaveRequest, error)
	Cancel(ctx context.Context, id uint, studentId uint) error
	Exists(ctx context.Context, id uint, studentId uint) (*bool, error)
}
