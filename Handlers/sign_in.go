package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	// TODO: Validate incoming request data.
	c.Status(http.StatusNoContent)
}
