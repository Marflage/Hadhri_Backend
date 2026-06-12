package dbmodels

import "time"

type AccountActivationRequest struct {
	Id              uint      `db:"id"`
	InsertedAt      time.Time `db:"inserted_at"`
	FullName        string    `db:"full_name"`
	Email           string    `db:"email"`
	PhoneNumber     string    `db:"phone_number"`
	PasswordHash    string    `db:"password_hash"`
	CoursePlanId    uint      `db:"course_plan_id"`
	Semester        uint      `db:"semester"`
	StatusChangedAt time.Time `db:"status_changed_at"`
}
