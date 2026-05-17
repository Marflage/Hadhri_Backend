package webapi

import (
	dtos "hadhri/Admin/Application/Dtos"
	usecases "hadhri/Admin/Application/UseCases"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getStudentRequest struct {
	Id int `form:"id" binding:"required"`
}

type getStudentHandler struct {
	uc usecases.GetStudent
}

func NewGetStudentHandler(uc usecases.GetStudent) getStudentHandler {
	return getStudentHandler{uc: uc}
}

func (h getStudentHandler) Handle(c *gin.Context) {
	var req getStudentRequest
	res := responses.ApiResponse[dtos.Student]{}

	if err := c.ShouldBindQuery(&req); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	dto, err := h.uc.Execute(c.Request.Context(), req.Id)

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Data = dto
	res.Message = "Successfully retrieved student."

	c.JSON(http.StatusOK, res)
}
