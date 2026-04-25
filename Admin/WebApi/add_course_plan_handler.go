package webapi

import (
	commands "hadhri/Admin/Application/Commands"
	usecases "hadhri/Admin/Application/UseCases"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addCoursePlanRequest struct {
	CourseId        int   `json:"courseId" binding:"required,gte=0"`
	ClassScheduleId int   `json:"classScheduleId" binding:"required,gte=0"`
	ClassSessionId  int   `json:"classSessionId" binding:"required,gte=0"`
	IsActive        *bool `json:"isActive" binding:"required"`
}

type AddCoursePlanHandler struct {
	uc usecases.AddCoursePlan
}

func NewAddCoursePlanHandler(uc usecases.AddCoursePlan) AddCoursePlanHandler {
	return AddCoursePlanHandler{uc: uc}
}

func (h AddCoursePlanHandler) AddCoursePlan(c *gin.Context) {
	var req addCoursePlanRequest
	res := responses.ApiResponse[any]{}

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// TODO: Validate request.

	cmd := commands.AddCoursePlan{
		CourseId:        req.CourseId,
		ClassScheduleId: req.ClassScheduleId,
		ClassSessionId:  req.ClassSessionId,
		IsActive:        *req.IsActive,
	}

	if err := h.uc.Execute(c.Request.Context(), cmd); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Successfully added course plan."

	c.JSON(http.StatusCreated, res)
}
