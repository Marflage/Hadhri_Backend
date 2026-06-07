package student

import (
	"fmt"
	domain "hadhri/StudentManagement/Domain"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// TODO: Should enrollment details be set through this aggregate root?
// TODO: Is there any way to hide the type from other packages? That is, to enforce calling the constructor.
type Student struct {
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	password    string
	enrollment  domain.Enrollment
}

func NewStudent(fullName string, email string, phoneNumber string, password string, coursePlanId int, semester int) (*Student, error) {
	student := Student{}

	if err := student.setFullname(fullName); err != nil {
		return nil, err
	}

	if err := student.setEmail(email); err != nil {
		return nil, err
	}

	if err := student.setPhoneNumber(phoneNumber); err != nil {
		return nil, err
	}

	if err := student.setPassword(password); err != nil {
		return nil, err
	}

	enrollment, err := domain.NewEnrollment(coursePlanId, semester)

	if err != nil {
		return nil, fmt.Errorf("Error creating enrollment: %w", err)
	}

	student.enrollment = enrollment

	return &student, nil
}

// Getters

func (s *Student) GetFullName() string {
	return s.firstName
}

func (s *Student) GetEmail() string {
	return s.email
}

func (s *Student) GetPhoneNumber() string {
	return s.phoneNumber
}

func (s *Student) GetPassword() string {
	return s.password
}

func (s *Student) GetCoursePlanId() int {
	return s.enrollment.CoursePlanId
}

func (s *Student) GetSemester() int {
	return s.enrollment.Semester
}

// Setters
// TODO: Sanitize input strings in all the methods.

func (s *Student) setFullname(value string) error {
	err := validateName(value)

	if err != nil {
		return fmt.Errorf("Invalid first name.")
	}

	s.firstName = value

	return nil
}

func (s *Student) setEmail(value string) error {
	err := validateEmail(value)

	if err != nil {
		return err
	}

	s.email = value

	return nil
}

func (s *Student) setPhoneNumber(value string) error {
	err := validatePhoneNumber(value)

	if err != nil {
		return err
	}

	s.phoneNumber = value

	return nil
}

func (s *Student) setPassword(value string) error {
	err := validatePassword(value)

	if err != nil {
		return err
	}

	// TODO: Move this into the Student entity to enforce hashing of password.
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)

	if err != nil {
		// TODO: Log.
		return err
	}

	s.password = string(passwordHash)

	return nil
}

// TODO: Attach these methods to the Student entity.
// TODO: Add constraint for maximum length.
func validateName(value string) error {
	if value == "" || len(value) < 3 {
		return fmt.Errorf("Invalid name.")
	}

	return nil
}

func validateEmail(value string) error {
	// TODO: Implement regex-based robust email validation.
	if value == "" || !strings.Contains(value, "@") {
		return fmt.Errorf("Invalid email.")
	}

	return nil
}

func validatePhoneNumber(value string) error {
	if len(value) != 11 {
		return fmt.Errorf("Invalid phone number.")
	}

	return nil
}

func validatePassword(value string) error {
	if value == "" || len(value) < 8 {
		return fmt.Errorf("Password must be 8 or more characters.")
	}

	return nil
}

// Errors

// type invalidStringError struct{}

// func (e invalidStringError) Error() string {
// 	return ""
// }

// type invalidFirstNameError struct {
// 	message string
// }

// func (e invalidFirstNameError) Error() string {
// 	return e.message
// }

// type invalidEmailError struct {
// 	// TODO: How to set default value for fields?
// 	message string
// }

// func (e invalidEmailError) Error() string {
// 	return e.message
// }
