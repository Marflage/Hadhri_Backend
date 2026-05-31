package handlers

import (
	commands "hadhri/LeaveManagement/Application/Commands"
	usecases "hadhri/LeaveManagement/Application/UseCases"
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

type pathParam struct {
	Id uint `uri:"id" binding:"required,gt=0"`
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

	var param pathParam

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

	cmd := commands.EditLeaveRequest{
		Id: id,
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
