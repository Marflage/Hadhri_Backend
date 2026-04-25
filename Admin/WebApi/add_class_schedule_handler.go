package webapi

import (
	commands "hadhri/Admin/Application/Commands"
	usecases "hadhri/Admin/Application/UseCases"
	responses "hadhri/WebApi/Responses"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type addClassScheduleRequest struct {
	Name string `json:"name" binding:"required"`
}

type classScheduleHandler struct {
	uc usecases.AddClassSchedule
}

func NewClassScheduleHandler(uc usecases.AddClassSchedule) classScheduleHandler {
	return classScheduleHandler{uc: uc}
}

func (h classScheduleHandler) AddClassSchedule(c *gin.Context) {
	var req addClassScheduleRequest
	res := responses.ApiResponse[any]{}

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	req.Name = strings.TrimSpace(req.Name)

	cmd := commands.AddClassSchedule{
		Name: req.Name,
	}

	if err := h.uc.Execute(c.Request.Context(), cmd); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Successfully added a class schedule."

	c.JSON(http.StatusCreated, res)
}
