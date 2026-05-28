package repositories

import (
	"context"
	domain "hadhri/LeaveManagement/Domain/LeaveRequest"
)

type ILeaveRequest interface {
	AddLeaveRequest(ctx context.Context, entity domain.LeaveRequest) error
	AlreadyExists(ctx context.Context, entity domain.LeaveRequest) error
}
