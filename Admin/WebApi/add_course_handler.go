package webapi

import (
	commands "hadhri/Admin/Application/Commands"
	usecases "hadhri/Admin/Application/UseCases"
	responses "hadhri/WebApi/Responses"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type addCourseRequest struct {
	Name string `json:"name" binding:"required"`
}

// TODO: Will this stay as a general struct?
type addCourseHandler struct {
	// TODO: Should the type be a pointer?
	uc usecases.AddCourse
}

func NewAddCourseHandler(uc usecases.AddCourse) addCourseHandler {
	return addCourseHandler{uc: uc}
}

func (h addCourseHandler) AddCourse(c *gin.Context) {
	var req addCourseRequest
	// TODO: Research for a better alternative/pattern for not having data in the response.
	res := responses.ApiResponse[any]{}

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	req.Name = strings.TrimSpace(req.Name)

	cmd := commands.AddCourse{
		Name: req.Name,
	}

	if err := h.uc.Execute(c.Request.Context(), cmd); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Successfully added a course."

	c.JSON(http.StatusCreated, res)
}
