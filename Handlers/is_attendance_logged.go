package handlers

import (
	"errors"
	db "hadhri/Db"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func IsAttendanceLogged(c *gin.Context) {
	res := responses.ApiResponse[bool]{}

	studentId := c.GetInt("studentId")

	dbConn, err := db.InitDb()

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	isAttendanceLogged, err := isAttendanceAlreadyLogged(studentId, dbConn)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			res.Error = err.Error()
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}
	}

	res.Data = isAttendanceLogged
	res.Message = "Fetched attendance status successfully."

	c.JSON(http.StatusOK, res)
}
