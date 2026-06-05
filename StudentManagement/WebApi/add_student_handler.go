package webapi

import (
	commands "hadhri/StudentManagement/Application/Commands"
	usecases "hadhri/StudentManagement/Application/Usecases"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Add the remaining binding tags.
type addStudentRequest struct {
	Fullname        string `json:"fullName" binding:"required"`
	Email           string `json:"email" binding:"required"`
	PhoneNumber     string `json:"phoneNumber" binding:"required"`
	CourseId        int    `json:"courseId" binding:"required"`
	ClassScheduleId int    `json:"classScheduleId" binding:"required"`
	ClassSessionId  int    `json:"classSessionId" binding:"required"`
	Semester        int    `json:"semester" binding:"required"`
}

type addStudentResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type addStudentHandler struct {
	uc usecases.AddStudent
}

func NewAddStudentHandler(uc usecases.AddStudent) addStudentHandler {
	return addStudentHandler{uc: uc}
}

func (h addStudentHandler) Handle(c *gin.Context) {
	var req addStudentRequest
	res := responses.ApiResponse[addStudentResponse]{}

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// TODO: Sanitize request strings.

	cmd := commands.AddStudent{
		FullName:        req.Fullname,
		Email:           req.Email,
		PhoneNumber:     req.PhoneNumber,
		CourseId:        req.CourseId,
		ClassScheduleId: req.ClassScheduleId,
		ClassSessionId:  req.ClassSessionId,
		Semester:        req.Semester,
	}

	creds, err := h.uc.Execute(c.Request.Context(), cmd)

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Data = addStudentResponse{
		Email:    creds.Email,
		Password: creds.Password,
	}
	res.Message = "Successfully added a student."

	c.JSON(http.StatusCreated, res)
}
