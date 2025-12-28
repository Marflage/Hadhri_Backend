package handlers

import (
	"context"
	db "hadhri/Db"
	entities "hadhri/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetCoursePlans(c *gin.Context) {
	conn, err := db.InitDb()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	sql := `
		SELECT *
		FROM course_plans
	`

	rows, err := conn.Query(context.Background(), sql)

	coursePlans, err := pgx.CollectRows(rows, pgx.RowToStructByName[entities.CoursePlan])

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(coursePlans) != 0 {
		c.JSON(http.StatusOK, coursePlans)
	}
}
