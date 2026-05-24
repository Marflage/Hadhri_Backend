package webapi

import (
	commands "hadhri/Auth/Application/Commands"
	usecases "hadhri/Auth/Application/Usecases"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signUpRequest struct {
	FirstName       string `json:"firstName" binding:"required,noBlank,min=2,max=30"`
	LastName        string `json:"lastName" binding:"required,noBlank,min=2,max=30"`
	Email           string `json:"email" binding:"required,email"`
	PhoneNumber     string `json:"phoneNumber" binding:"required,numeric,len=11"`
	Password        string `json:"password" binding:"required,min=8,max=30,noBlank"`
	CourseId        int    `json:"courseId" binding:"required,gte=1"`
	ClassScheduleId int    `json:"classScheduleId" binding:"required,gte=1"`
	ClassSessionId  int    `json:"classSessionId" binding:"required,gte=1"`
	// TODO: Create different enums for semester number for each course
	Semester int `json:"semester" binding:"required,min=1,max=8"`
}

type signUpHandler struct {
	uc usecases.SignUp
}

func NewSignUpHandler(cs usecases.SignUp) signUpHandler {
	return signUpHandler{uc: cs}
}

func (h signUpHandler) Handle(c *gin.Context) {
	var req signUpRequest
	res := responses.ApiResponse[responses.Auth]{}

	if err := c.ShouldBind(&req); err != nil {
		// TODO: Log error
		// TODO: Create a util to make the error messages more readable and return that.
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	cmd := commands.SignUp{
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		Email:           req.Email,
		PhoneNumber:     req.PhoneNumber,
		Password:        req.Password,
		CourseId:        req.CourseId,
		ClassScheduleId: req.ClassScheduleId,
		ClassSessionId:  req.ClassSessionId,
		Semester:        req.Semester,
	}

	tokenPtr, err := h.uc.Execute(c.Request.Context(), cmd)

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	// res.Data.StudentId = studentId
	res.Data.Token = *tokenPtr
	res.Message = "Signed up successfully."

	c.JSON(http.StatusOK, res)
}
