package handlers

import (
	"context"
	"database/sql"
	"errors"
	db "hadhri/Db"
	querymodels "hadhri/QueryModels"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetStudentEnrollmentDetails(c *gin.Context) {
	res := responses.ApiResponse[*responses.GetStudentEnrollmentDetails]{}

	studentId := c.GetInt("studentId")

	dbConn, err := db.InitDb()

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	// TODO: Why is not the address of dbConn required to be sent here as the callee has a pointer parameter?
	enrollmentDetails, err := getStudentEnrollmentDetails(dbConn, studentId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res.Error = "No data found."
		} else {
			res.Error = err.Error()
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	data := mapQueryModelToResponseDto2(*enrollmentDetails)

	res.Data = &data
	res.Message = "Successfully fetched student enrollment data."

	c.IndentedJSON(http.StatusOK, res)
}

func getStudentEnrollmentDetails(dbConn *pgx.Conn, studentId int) (*querymodels.GetStudentEnrollmentDetails, error) {
	query := `
	SELECT c.name       AS CourseName,
       cs.name      AS ClassScheduleName,
       s.name       AS ClassSessionName,
       s.start_time AS ClassSessionStartTime,
       s.end_time   AS ClassSessionEndTime,
       se.semester  AS Semester
	FROM enrollments se
			INNER JOIN public.course_plans cp
						ON cp.id = se.course_plan_id
			INNER JOIN public.courses c
						ON c.id = cp.course_id
			INNER JOIN public.class_schedules cs
						ON cs.id = cp.class_schedule_id
			INNER JOIN public.class_sessions s
						ON s.id = cp.class_session_id
	WHERE se.student_id = $1
	`

	var result querymodels.GetStudentEnrollmentDetails

	if err := dbConn.QueryRow(context.Background(), query, studentId).Scan(&result.CourseName, &result.ClassScheduleName, &result.ClassSessionName, &result.ClassSessionStartTime, &result.ClassSessionEndTime, &result.Semester); err != nil {
		return nil, err
	}

	return &result, nil
}

func mapQueryModelToResponseDto2(queryModel querymodels.GetStudentEnrollmentDetails) responses.GetStudentEnrollmentDetails {
	dto := responses.GetStudentEnrollmentDetails{
		CourseName:            queryModel.CourseName,
		ClassScheduleName:     queryModel.ClassScheduleName,
		ClassSessionName:      queryModel.ClassSessionName,
		ClassSessionStartTime: queryModel.ClassSessionStartTime,
		ClassSessionEndTime:   queryModel.ClassSessionEndTime,
		Semester:              queryModel.Semester,
	}

	return dto
}
