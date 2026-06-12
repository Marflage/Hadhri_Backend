package handlers

import (
	usecases "hadhri/StudentManagement/Application/Usecases"
	webapi "hadhri/StudentManagement/WebApi"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type approveAccountActivationHandler struct {
	uc usecases.ApproveAccountActivation
}

func NewApproveAccountActivationHandler(uc usecases.ApproveAccountActivation) approveAccountActivationHandler {
	return approveAccountActivationHandler{uc: uc}
}

func (self approveAccountActivationHandler) Handle(c *gin.Context) {
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

	res.Message = "Approved account actvation request successfully."

	c.JSON(http.StatusOK, res)
}
