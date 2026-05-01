package webapi

import (
	usecases "hadhri/Admin/Application/UseCases"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type classSchedule struct {
	Name string
}

type getAllClassSchedulesHandler struct {
	uc usecases.GetAllClassSchedules
}

func NewGetAllClassSchedulesHandler(uc usecases.GetAllClassSchedules) getAllClassSchedulesHandler {
	return getAllClassSchedulesHandler{uc: uc}
}

func (h getAllClassSchedulesHandler) GetAll(c *gin.Context) {
	res := responses.ApiResponse[[]classSchedule]{}

	dtos, err := h.uc.Execute(c.Request.Context())

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	var classSchedules []classSchedule

	for _, dto := range dtos {
		classSchedule := classSchedule{
			Name: dto.Name,
		}

		classSchedules = append(classSchedules, classSchedule)
	}

	res.Data = classSchedules
	res.Message = "Successfully retrieved class schedules."

	c.JSON(http.StatusOK, res)
}
