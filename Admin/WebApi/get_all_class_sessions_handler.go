package webapi

import (
	usecases "hadhri/Admin/Application/UseCases"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type classSession struct {
	Name string `json:"name"`
}

type getAllClassSessionsHandler struct {
	uc usecases.GetAllClassSessions
}

func NewGetAllClassSessionsHandler(uc usecases.GetAllClassSessions) getAllClassSessionsHandler {
	return getAllClassSessionsHandler{uc: uc}
}

func (h getAllClassSessionsHandler) GetAll(c *gin.Context) {
	res := responses.ApiResponse[[]classSession]{}

	dtos, err := h.uc.Execute(c.Request.Context())

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	var classSessions []classSession

	for _, dto := range dtos {
		classSession := classSession{
			Name: dto.Name,
		}

		classSessions = append(classSessions, classSession)
	}

	res.Data = classSessions
	res.Message = "Successfully retrieved class sessions."

	c.JSON(http.StatusOK, res)
}
