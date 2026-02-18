package dtos

type ClassSession struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	AvailableSemesters []int  `json:"availableSemesters"`
}
