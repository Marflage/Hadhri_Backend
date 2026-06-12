package handlers

import (
	usecases "hadhri/StudentManagement/Application/Usecases"
	webapi "hadhri/StudentManagement/WebApi"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type declineAccountActivationHandler struct {
	uc usecases.DeclineAccountActivation
}

func NewDeclineAccountActivationHandler(uc usecases.DeclineAccountActivation) declineAccountActivationHandler {
	return declineAccountActivationHandler{uc: uc}
}

func (self declineAccountActivationHandler) Handle(c *gin.Context) {
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

	res.Message = "Declined account activation request successfully."

	c.JSON(http.StatusOK, res)
}
