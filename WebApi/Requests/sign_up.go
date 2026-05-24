package requests

import (
	"encoding/json"
	"fmt"
)

type CourseName string

// TODO: Find a way to synchronize reference data in the DB and the values of the enum.
const (
	QS  CourseName = "Quranic Sciences (QS)"
	QHD CourseName = "Quran & Hadeeth Dimensions (QHD)"
	QR  CourseName = "Quranic Reflections (QR)"
	DEN CourseName = "Dars e Nizami"
)

func (courseName *CourseName) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch CourseName(s) {
	case QS, QHD, QR, DEN:
		*courseName = CourseName(s)
		return nil
	default:
		return fmt.Errorf("Invalid course name: %s", s)
	}
}

type ClassSchedule string

// TODO: Find a way to synchronize reference data in the DB and the values of the enum.
const (
	Weekday ClassSchedule = "Weekday"
	Weekend ClassSchedule = "Weekend"
)

func (classSchedule *ClassSchedule) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch ClassSchedule(s) {
	case Weekday, Weekend:
		*classSchedule = ClassSchedule(s)
		return nil
	default:
		return fmt.Errorf("Invalid class schedule: %s", s)
	}
}

type ClassSession string

// TODO: Find a way to synchronize reference data in the DB and the values of the enum.
// TODO: Include time along with session name
const (
	Morning   ClassSession = "Morning"
	Afternoon ClassSession = "Afternoon"
	Evening   ClassSession = "Evening"
)

func (classSession *ClassSession) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch ClassSession(s) {
	case Morning, Afternoon, Evening:
		*classSession = ClassSession(s)
		return nil
	default:
		return fmt.Errorf("Invalid class session: %s", s)
	}
}
