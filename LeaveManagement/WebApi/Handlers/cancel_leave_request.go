package handlers

import (
	usecases "hadhri/LeaveManagement/Application/UseCases"
	webapi "hadhri/LeaveManagement/WebApi"
	constants "hadhri/LeaveManagement/WebApi/Constants"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cancelLeaveRequestHandler struct {
	uc usecases.CancelLeaveRequest
}

func NewCancelLeaveRequestHandler(uc usecases.CancelLeaveRequest) cancelLeaveRequestHandler {
	return cancelLeaveRequestHandler{uc: uc}
}

func (self cancelLeaveRequestHandler) Handle(c *gin.Context) {
	res := responses.ApiResponse[any]{}

	var pathParam webapi.PathParam

	if err := c.ShouldBindUri(&pathParam); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	id := pathParam.Id
	studentId := c.GetUint(constants.StudentId)

	if err := self.uc.Execute(c.Request.Context(), id, studentId); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Leave request canceled successfully."

	c.JSON(http.StatusOK, res)
}
