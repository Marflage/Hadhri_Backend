package handlers

import (
	"context"
	"database/sql"
	"errors"
	db "hadhri/Db"
	querymodels "hadhri/QueryModels"
	requests "hadhri/Requests"
	responses "hadhri/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetStudentDetails(c *gin.Context) {
	var req requests.GetStudentDetails
	// TODO: Find a better alternative to make the data nullable.
	res := &responses.ApiResponse[*responses.GetStudentDetails]{}

	if err := c.ShouldBindQuery(&req); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// TODO: Find a way to inject this into all the handlers instead.
	dbConn, err := db.InitDb()

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	studentDetails, err := getStudentDetails(dbConn, req.StudentId)

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
	SELECT s.id         AS StudentId,
       s.first_name AS FirstName,
       s.last_name  AS LastName,
       c.name       AS CourseName,
       cs.name      AS ClassScheduleName,
       cses.name    AS ClassSessionName,
       se.semester  AS Semester
	FROM students s
			INNER JOIN public.student_enrollments se
						ON s.id = se.student_id
			INNER JOIN public.course_plans cp
						ON se.course_plan_id = cp.id
			INNER JOIN public.courses c
						ON cp.course_id = c.id
			INNER JOIN public.class_schedules cs
						ON cp.class_schedule_id = cs.id
			INNER JOIN public.class_sessions cses
						ON cp.class_session_id = cses.id
	WHERE s.id = $1
	`

	var result querymodels.GetStudentDetails

	if err := dbConn.QueryRow(context.Background(), query, id).
		// TODO: How to map the values directly to a struct?
		Scan(&result.StudentId, &result.FirstName, &result.LastName, &result.CourseName, &result.ClassScheduleName, &result.ClassSessionName, &result.Semester); err != nil {
		return nil, err
	}

	return &result, nil
}

func mapQueryModelToResponseDto(queryModel querymodels.GetStudentDetails) responses.GetStudentDetails {
	dto := responses.GetStudentDetails{
		StudentId:         queryModel.StudentId,
		FirstName:         queryModel.FirstName,
		LastName:          queryModel.LastName,
		CourseName:        queryModel.CourseName,
		ClassScheduleName: queryModel.ClassScheduleName,
		ClassSessionName:  queryModel.ClassSessionName,
		Semester:          queryModel.Semester,
	}

	return dto
}
