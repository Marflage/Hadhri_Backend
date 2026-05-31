package handlers

import (
	commands "hadhri/LeaveManagement/Application/Commands"
	usecases "hadhri/LeaveManagement/Application/UseCases"
	webapi "hadhri/LeaveManagement/WebApi"
	constants "hadhri/LeaveManagement/WebApi/Constants"
	responses "hadhri/WebApi/Responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type editLeaveRequestRequest struct {
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
	Reason    *string    `json:"reason"`
}

type editLeaveRequestHandler struct {
	uc usecases.EditLeaveRequest
}

func NewEditLeaveHandler(uc usecases.EditLeaveRequest) editLeaveRequestHandler {
	return editLeaveRequestHandler{uc: uc}
}

func (self editLeaveRequestHandler) Handle(c *gin.Context) {
	var req editLeaveRequestRequest
	res := responses.ApiResponse[any]{}

	var param webapi.PathParam

	if err := c.ShouldBindUri(&param); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	id := param.Id

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// TODO: Even though validation will be identical to domain entity validation, should requests be validated anyway?

	studentId := c.GetUint(constants.StudentId)

	cmd := commands.EditLeaveRequest{
		Id:        id,
		StudentId: studentId,
	}

	if req.StartDate != nil {
		cmd.StartDate = req.StartDate
	}

	if req.EndDate != nil {
		cmd.EndDate = req.EndDate
	}

	if req.Reason != nil {
		cmd.Reason = req.Reason
	}

	if err := self.uc.Execute(c.Request.Context(), cmd); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Leave request edited successfully."

	c.JSON(http.StatusOK, res)
}
