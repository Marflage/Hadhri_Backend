package webapi

import (
	commands "hadhri/Admin/Application/Commands"
	usecases "hadhri/Admin/Application/UseCases"
	responses "hadhri/WebApi/Responses"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type addClassSessionRequest struct {
	Name      string    `json:"name" binding:"required"`
	StartTime time.Time `json:"startTime" binding:"required"`
	EndTime   time.Time `json:"endTime" binding:"required"`
}

// TODO: Rename this to just ClassSessionHandler as more methods will be added for new endpoints.
type addClassSessionHandler struct {
	uc usecases.AddClassSession
}

func NewAddClassSessionHandler(uc usecases.AddClassSession) addClassSessionHandler {
	return addClassSessionHandler{uc: uc}
}

func (h addClassSessionHandler) AddClassSession(c *gin.Context) {
	var req addClassSessionRequest
	res := responses.ApiResponse[any]{}

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// TODO: Validate if endTime > startTime.
	// TODo: Is this validation even necessary as the use case will validate anyway?
	// TODO: Can this validation be done using the binding tags?
	if !req.EndTime.After(req.StartTime) {
		res.Error = "End time must be after start time."
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	req.Name = strings.TrimSpace(req.Name)

	cmd := commands.AddClassSession{
		Name:      req.Name,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	if err := h.uc.Execute(c.Request.Context(), cmd); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Message = "Successfully added class session."

	c.JSON(http.StatusCreated, res)
}
