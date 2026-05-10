package webapi

import (
	usecases "hadhri/Admin/Application/UseCases"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type coursePlanResponse struct {
	Id                int    `json:"id"`
	CourseId          int    `json:"courseId"`
	CourseName        string `json:"courseName"`
	ClassScheduleId   int    `json:"classScheduleId"`
	ClassScheduleName string `json:"classScheduleName"`
	ClassSessionId    int    `json:"classSessionId"`
	ClassSessionName  string `json:"classSessionName"`
	IsActive          bool   `json:"isActive"`
}

type getAllCoursePlansHandler struct {
	uc usecases.GetAllCoursePlans
}

func NewGetAllCoursePlansHandler(uc usecases.GetAllCoursePlans) getAllCoursePlansHandler {
	return getAllCoursePlansHandler{uc: uc}
}

func (h getAllCoursePlansHandler) GetAll(c *gin.Context) {
	res := responses.ApiResponse[[]coursePlanResponse]{}

	dtos, err := h.uc.Execute(c.Request.Context())

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	var coursePlans []coursePlanResponse

	for _, dto := range dtos {
		coursePlan := coursePlanResponse{
			Id:                dto.Id,
			CourseId:          dto.CourseId,
			CourseName:        dto.CourseName,
			ClassScheduleId:   dto.ClassScheduleId,
			ClassScheduleName: dto.ClassScheduleName,
			ClassSessionId:    dto.ClassSessionId,
			ClassSessionName:  dto.ClassSessionName,
			IsActive:          dto.IsActive,
		}

		coursePlans = append(coursePlans, coursePlan)
	}

	res.Data = coursePlans
	res.Message = "Successfully retrieved all course plans."

	c.JSON(http.StatusOK, res)
}
