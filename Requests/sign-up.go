package requests

// TODO: Is not there a better way to bind data instead of using string tags?
type SignUpRequest struct {
	FirstName   string     `form:"firstName" binding:"required"`
	LastName    string     `form:"lastName" binding:"required"`
	Email       string     `form:"email" binding:"required,email"`
	PhoneNumber string     `form:"phoneNumber" binding:"required,max=11"`
	CourseName  CourseName `form:"courseName" binding:"required"`
	// TODO: Create different enums for semester number for each course
	Semester      string        `form:"semester" binding:"required,min=1"`
	ClassSchedule ClassSchedule `form:"classSchedule" binding:"required"`
	ClassSession  ClassSession  `form:"classSession" binding:"required"`
	// TimeSlot      TimeSlot      `form:"timeSlot"`
	Password string `form:"password" binding:"required"`
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
