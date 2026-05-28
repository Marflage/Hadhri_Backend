package domain

import (
	"fmt"
	"time"
)

type LeaveRequest struct {
	id int
	// TODO: Should this be a reference to the ID in the Student entity?
	studentId int
	startDate time.Time
	endDate   time.Time
	reason    string
}

func NewLeaveRequest(studentId int, startDate time.Time, endDate time.Time, reason string) (*LeaveRequest, error) {
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

	return &e, nil
}

// Getters

func (e LeaveRequest) GetId() int {
	return e.id
}

func (e LeaveRequest) GetStudentId() int {
	return e.studentId
}

func (e LeaveRequest) GetStartDate() time.Time {
	return e.startDate
}

func (e LeaveRequest) GetEndDate() time.Time {
	return e.endDate
}

func (e LeaveRequest) GetReason() string {
	return e.reason
}

// Setters

func (e *LeaveRequest) setStudentId(id int) error {
	if err := e.validateStudentId(id); err != nil {
		return err
	}

	e.studentId = id
	return nil
}

func (e *LeaveRequest) setStartDate(startDate time.Time) error {
	if err := e.validateStartDate(startDate); err != nil {
		return err
	}

	e.startDate = startDate
	return nil
}

func (e *LeaveRequest) setEndDate(endDate time.Time) error {
	if err := e.validateEndDate(endDate); err != nil {
		return err
	}

	e.endDate = endDate
	return nil
}

func (e *LeaveRequest) setReason(reason string) error {
	if err := e.validateReason(reason); err != nil {
		return err
	}

	e.reason = reason
	return nil
}

// Validators

func (e LeaveRequest) validateStudentId(id int) error {
	if id <= 0 {
		return fmt.Errorf("Invalid student ID.")
	}

	return nil
}

func (e *LeaveRequest) validateStartDate(startDate time.Time) error {
	today := time.Now().Truncate(24 * time.Hour)

	if startDate.Before(today) {
		return fmt.Errorf("Invalid start date.")
	}

	return nil
}

func (e LeaveRequest) validateEndDate(endDate time.Time) error {
	today := time.Now().Truncate(24 * time.Hour)

	if endDate.Before(today) {
		return fmt.Errorf("Invalid end date.")
	}

	return nil
}

func (e LeaveRequest) validateReason(reason string) error {
	if reason == "" || len(reason) < 4 {
		return fmt.Errorf("Invalid reason.")
	}

	return nil
}

func (e LeaveRequest) validateStartDateAndEndDate() error {
	if e.startDate.After(e.endDate) {
		return fmt.Errorf("Invalid start date.")
	}

	if e.endDate.Before(e.startDate) {
		return fmt.Errorf("Invalid end date.")
	}

	return nil
}
