package domain

import (
	"errors"
	"fmt"
	"time"
)

type LeaveRequest struct {
	id uint
	// TODO: Should this be a reference to the ID in the Student entity?
	studentId uint
	startDate time.Time
	endDate   time.Time
	reason    string
	// TODO: Use an enum.
	status string
}

// For creating a new entity
func NewLeaveRequest(studentId uint, startDate time.Time, endDate time.Time, reason string) (*LeaveRequest, error) {
	e := LeaveRequest{}

	if err := e.setStudentId(studentId); err != nil {
		return nil, err
	}

	if err := e.setStartDate(startDate); err != nil {
		return nil, err
	}

	if err := e.setEndDate(endDate); err != nil {
		return nil, err
	}

	e.validateStartDateAndEndDate()

	if err := e.setReason(reason); err != nil {
		return nil, err
	}

	e.status = "pending"

	return &e, nil
}

// For loading an exisitng entity
func ReconstituteLeaveRequest(id uint, studentId uint, startDate time.Time, endDate time.Time, reason string, status string) LeaveRequest {
	e := LeaveRequest{
		id:        id,
		studentId: studentId,
		startDate: startDate,
		endDate:   endDate,
		reason:    reason,
		status:    status,
	}

	return e
}

// Getters

func (self LeaveRequest) GetId() uint {
	return self.id
}

func (self LeaveRequest) GetStudentId() uint {
	return self.studentId
}

func (self LeaveRequest) GetStartDate() time.Time {
	return self.startDate
}

func (self LeaveRequest) GetEndDate() time.Time {
	return self.endDate
}

func (self LeaveRequest) GetReason() string {
	return self.reason
}

func (self LeaveRequest) GetStatus() string {
	return self.status
}

// Behavior methods

func (self *LeaveRequest) Reschedule(startDate *time.Time, endDate *time.Time) error {
	if startDate != nil {
		if err := self.setStartDate(*startDate); err != nil {
			return err
		}
	}

	if endDate != nil {
		if err := self.setEndDate(*endDate); err != nil {
			return err
		}
	}

	if err := self.validateStartDateAndEndDate(); err != nil {
		return err
	}

	return nil
}

func (self *LeaveRequest) ChangeReason(reason string) error {
	return self.setReason(reason)
}

// Setters

func (self *LeaveRequest) setStudentId(id uint) error {
	if err := self.validateStudentId(id); err != nil {
		return err
	}

	self.studentId = id
	return nil
}

func (self *LeaveRequest) setStartDate(startDate time.Time) error {
	if err := self.validateStartDate(startDate); err != nil {
		return err
	}

	self.startDate = startDate
	return nil
}

func (self *LeaveRequest) setEndDate(endDate time.Time) error {
	if err := self.validateEndDate(endDate); err != nil {
		return err
	}

	self.endDate = endDate
	return nil
}

func (self *LeaveRequest) setReason(reason string) error {
	if err := self.validateReason(reason); err != nil {
		return err
	}

	self.reason = reason
	return nil
}

func (self *LeaveRequest) setStatus(status string) error {
	if err := self.validateStatus(status); err != nil {
		return err
	}

	self.status = status
	return nil
}

// Validators

func (self LeaveRequest) validateStudentId(id uint) error {
	if id <= 0 {
		return fmt.Errorf("Invalid student ID.")
	}

	return nil
}

func (self *LeaveRequest) validateStartDate(startDate time.Time) error {
	today := time.Now().Truncate(24 * time.Hour)

	if startDate.Before(today) {
		return fmt.Errorf("Invalid start date.")
	}

	return nil
}

func (self LeaveRequest) validateEndDate(endDate time.Time) error {
	today := time.Now().Truncate(24 * time.Hour)

	if endDate.Before(today) {
		return fmt.Errorf("Invalid end date.")
	}

	return nil
}

func (self LeaveRequest) validateReason(reason string) error {
	if reason == "" || len(reason) < 4 {
		return fmt.Errorf("Invalid reason.")
	}

	return nil
}

func (self LeaveRequest) validateStartDateAndEndDate() error {
	if self.startDate.After(self.endDate) {
		return fmt.Errorf("Start date cannot be after end date.")
	}

	if self.endDate.Before(self.startDate) {
		return fmt.Errorf("End date cannot be before start date.")
	}

	return nil
}

func (self LeaveRequest) validateStatus(status string) error {
	// TODO: Add valid transition rules, such as approved or rejected cannot be transitioned to pending again. Define all the valid transition paths.
	switch status {
	case "pending", "approved", "rejected":
		return nil
	default:
		return errors.New("Invalid status.")
	}
}
