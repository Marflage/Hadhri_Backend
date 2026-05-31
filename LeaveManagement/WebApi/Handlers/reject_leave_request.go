package handlers

import (
	usecases "hadhri/LeaveManagement/Application/UseCases"
	webapi "hadhri/LeaveManagement/WebApi"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type rejectLeaveRequestHandler struct {
	uc usecases.RejectLeaveRequest
}

func NewRejectLeaveRequest(uc usecases.RejectLeaveRequest) rejectLeaveRequestHandler {
	return rejectLeaveRequestHandler{uc: uc}
}

func (self rejectLeaveRequestHandler) Handle(c *gin.Context) {
	res := responses.ApiResponse[any]{}

	var param webapi.PathParam

	if err := c.ShouldBindUri(&param); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	id := param.Id

	if err := self.uc.Execute(c.Request.Context(), id); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Rejected leave request successfully."

	c.JSON(http.StatusOK, res)
}
