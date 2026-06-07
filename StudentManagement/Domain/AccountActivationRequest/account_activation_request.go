package accountactivationrequest

import (
	"errors"
	domain "hadhri/StudentManagement/Domain"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AccountActivationRequest struct {
	status          AccountActivationRequestStatus
	statusChangedAt time.Time
	fullName        string
	email           string
	phoneNumber     string
	passwordHash    string
	enrollment      domain.Enrollment
}

type AccountActivationRequestStatus string

const (
	pending  AccountActivationRequestStatus = "pending"
	approved AccountActivationRequestStatus = "approved"
	declined AccountActivationRequestStatus = "declined"
)

func NewAccountActivationRequest(fullname string,
	email string,
	phoneNumber string,
	password string,
	coursePlanId int,
	semester int) (*AccountActivationRequest, error) {
	if err := validateFullName(fullname); err != nil {
		return nil, err
	}

	if err := validateEmail(email); err != nil {
		return nil, err
	}

	if err := validatePhoneNumber(phoneNumber); err != nil {
		return nil, err
	}

	if err := validatePassword(password); err != nil {
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	enrollment, err := domain.NewEnrollment(coursePlanId, semester)

	if err != nil {
		return nil, err
	}

	return &AccountActivationRequest{
		status:          pending,
		statusChangedAt: time.Now(),
		fullName:        fullname,
		email:           email,
		phoneNumber:     phoneNumber,
		passwordHash:    string(passwordHash),
		enrollment:      enrollment,
	}, nil
}

// Getters

func (self AccountActivationRequest) Status() AccountActivationRequestStatus { return self.status }
func (self AccountActivationRequest) StatusChangedAt() time.Time             { return self.statusChangedAt }
func (self AccountActivationRequest) FullName() string                       { return self.fullName }
func (self AccountActivationRequest) Email() string                          { return self.email }
func (self AccountActivationRequest) PhoneNumber() string                    { return self.phoneNumber }
func (self AccountActivationRequest) PasswordHash() string                   { return self.passwordHash }
func (self AccountActivationRequest) CoursePlanId() int                      { return self.enrollment.CoursePlanId }
func (self AccountActivationRequest) Semester() int                          { return self.enrollment.Semester }

// Behaviors

func (self *AccountActivationRequest) Approve() error {
	if self.status != pending {
		return errors.New("Only pending requests can be approved.")
	}

	self.status = approved
	self.statusChangedAt = time.Now()

	return nil
}

func (self *AccountActivationRequest) Decline() error {
	if self.status != pending {
		return errors.New("Only pending requests can be declined.")
	}

	self.status = declined
	self.statusChangedAt = time.Now()

	return nil
}

// Validators

func validateFullName(value string) error {
	if len(strings.TrimSpace(value)) < 3 {
		return errors.New("Name must be at least 3 characters.")
	}

	return nil
}

func validateEmail(value string) error {
	// TODO: Implement regex-based robust email validation.
	if len(value) < 3 || !strings.Contains(value, "@") {
		return errors.New("Invalid email.")
	}

	return nil
}

func validatePhoneNumber(value string) error {
	if len(value) != 11 {
		return errors.New("Invalid phone number.")
	}

	for _, e := range value {
		if e < '0' || e > '9' {
			return errors.New("Phone number must contain only digits.")
		}
	}

	return nil
}

func validatePassword(value string) error {
	if len(value) < 8 {
		return errors.New("Password must be at least 8 characters.")
	}

	return nil
}
