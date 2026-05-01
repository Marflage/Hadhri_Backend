package webapi

import (
	usecases "hadhri/Admin/Application/UseCases"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type course struct {
	Name string `json:"name"`
}

type getAllCoursesHandler struct {
	uc usecases.GetAllCourses
}

func NewGetAllCoursesHandler(uc usecases.GetAllCourses) getAllCoursesHandler {
	return getAllCoursesHandler{uc: uc}
}

func (h getAllCoursesHandler) GetAll(c *gin.Context) {
	res := responses.ApiResponse[[]course]{}

	dtos, err := h.uc.Execute(c.Request.Context())

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	var courses []course

	for _, dto := range dtos {
		course := course{
			Name: dto.Name,
		}

		courses = append(courses, course)
	}

	res.Data = courses
	res.Message = "Successfully retrieved courses."

	c.JSON(http.StatusOK, res)
}
