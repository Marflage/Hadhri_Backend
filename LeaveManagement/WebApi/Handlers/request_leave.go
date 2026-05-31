package handlers

import (
	commands "hadhri/LeaveManagement/Application/Commands"
	usecases "hadhri/LeaveManagement/Application/UseCases"
	constants "hadhri/LeaveManagement/WebApi/Handlers/Constants"
	responses "hadhri/WebApi/Responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type requestLeaveRequest struct {
	StartDate time.Time `json:"startDate" binding:"required"`
	EndDate   time.Time `json:"endDate" binding:"required"`
	Reason    string    `json:"reason" binding:"required,min=4,max=200"`
}

type requestLeaveHandler struct {
	uc usecases.RequestLeave
}

func NewRequestLeaveHandler(uc usecases.RequestLeave) requestLeaveHandler {
	return requestLeaveHandler{uc: uc}
}

func (h requestLeaveHandler) Handle(c *gin.Context) {
	var req requestLeaveRequest
	res := responses.ApiResponse[any]{}

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	studentId := c.GetUint(constants.StudentId)

	cmd := commands.RequestLeave{
		StudentId: studentId,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Reason:    req.Reason,
	}

	if err := h.uc.Execute(c.Request.Context(), cmd); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Leave requested successfully."

	c.JSON(http.StatusCreated, res)
}
