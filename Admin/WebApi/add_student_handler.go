package webapi

import (
	commands "hadhri/Admin/Application/Commands"
	usecases "hadhri/Admin/Application/UseCases"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addStudentRequest struct {
	// TODO: Add validation binding tags.
	FirstName       string `json:"firstName" binding:"required"`
	LastName        string `json:"lastName" binding:"required"`
	Email           string `json:"email" binding:"required"`
	PhoneNumber     string `json:"phoneNumber" binding:"required"`
	Password        string `json:"password" binding:"required"`
	CourseId        int    `json:"courseId" binding:"required"`
	ClassScheduleId int    `json:"classScheduleId" binding:"required"`
	ClassSessionId  int    `json:"classSessionId" binding:"required"`
	Semester        int    `json:"semester" binding:"required"`
}

type addStudentHandler struct {
	uc usecases.AddStudent
}

func NewAddStudentHandler(uc usecases.AddStudent) addStudentHandler {
	return addStudentHandler{uc: uc}
}

func (h addStudentHandler) Handle(c *gin.Context) {
	var req addStudentRequest
	res := responses.ApiResponse[any]{}

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// TODO: Sanitize request strings.

	cmd := commands.AddStudent{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}

	if err := h.uc.Execute(c.Request.Context(), cmd); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	// TODO: Return the email and randomly-generated password.

	res.Message = "Successfully added a student."

	c.JSON(http.StatusCreated, res)
}
