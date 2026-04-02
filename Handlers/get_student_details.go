package handlers

import (
	"context"
	"database/sql"
	"errors"
	db "hadhri/Db"
	querymodels "hadhri/QueryModels"
	responses "hadhri/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetStudentDetails(c *gin.Context) {
	// TODO: Find a better alternative to make the data nullable.
	res := responses.ApiResponse[*responses.GetStudentDetails]{}

	studentId := c.GetInt("studentId")

	// TODO: Find a way to inject this into all the handlers instead.
	dbConn, err := db.InitDb()

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	studentDetails, err := getStudentDetails(dbConn, studentId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res.Error = "No data found."
			res.Data = nil
		} else {
			res.Error = err.Error()
		}

		// TODO: Change the status code to something more specific.
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	data := mapQueryModelToResponseDto(*studentDetails)

	res.Data = &data
	res.Message = "Successfully fetched student data."

	c.IndentedJSON(http.StatusOK, res)
}

func getStudentDetails(dbConn *pgx.Conn, id int) (*querymodels.GetStudentDetails, error) {
	query := `
	SELECT s.first_name AS FirstName,
		s.last_name  AS LastName
	FROM students s
	WHERE s.id = $1
	`

	var result querymodels.GetStudentDetails

	if err := dbConn.QueryRow(context.Background(), query, id).
		// TODO: How to map the values directly to a struct?
		Scan(&result.FirstName, &result.LastName); err != nil {
		return nil, err
	}

	return &result, nil
}

func mapQueryModelToResponseDto(queryModel querymodels.GetStudentDetails) responses.GetStudentDetails {
	dto := responses.GetStudentDetails{
		FirstName: queryModel.FirstName,
		LastName:  queryModel.LastName,
	}

	return dto
}
