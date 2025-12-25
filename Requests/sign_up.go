package requests

// TODO: Is not there a better way to bind data instead of using string tags?
type SignUpRequest struct {
	// TODO: Handle validation for API client-sent requests.
	// TODO: Create a custom validator for whitespace-only strings.
	FirstName   string     `form:"firstName" binding:"required,noBlank"`
	LastName    string     `form:"lastName" binding:"required,noBlank"`
	Email       string     `form:"email" binding:"required,email"`
	PhoneNumber string     `form:"phoneNumber" binding:"required,max=11,noBlank"`
	CourseName  CourseName `form:"courseName" binding:"required"`
	// TODO: Create different enums for semester number for each course
	Semester      int           `form:"semester" binding:"required,min=1,max=8"`
	ClassSchedule ClassSchedule `form:"classSchedule" binding:"required"`
	ClassSession  ClassSession  `form:"classSession" binding:"required"`
	Password      string        `form:"password" binding:"required,min=8,noBlank"`
}

type CourseName string

const (
	QS  CourseName = "Quranic Sciences (QS)"
	QHD CourseName = "Quran & Hadeeth Dimensions (QHD)"
	QR  CourseName = "Quranic Reflection (QR)"
	DEN CourseName = "Dars e Nizami"
)

type ClassSchedule string

const (
	Weekend ClassSchedule = "Weekend"
	weekday ClassSchedule = "Weekday"
)

type ClassSession string

// TODO: Include time along with session name
const (
	Morning   ClassSession = "Morning"
	Afternoon ClassSession = "Afternoon"
	Evening   ClassSession = "Evening"
)

// TODO: Create IsValid and Un/Marshal methods for all the enums
