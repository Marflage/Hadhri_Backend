package handlers

import (
	"context"
	"database/sql"
	"errors"
	db "hadhri/Db"
	requests "hadhri/Requests"
	responses "hadhri/Responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func LogAttendance(c *gin.Context) {
	var req requests.LogAttendance
	res := &responses.ApiResponse{}

	if err := c.ShouldBindQuery(&req); err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	dbConn, err := db.InitDb()

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	query := `
	SELECT ses.start_time, ses.end_time, cp.id
	FROM student_enrollments se
	INNER JOIN course_plans cp
	ON cp.id = se.course_plan_id
	INNER JOIN class_sessions ses
	ON ses.id = cp.class_session_id
	WHERE se.student_id = $1
	`

	var startTime time.Time
	var endTime time.Time
	var coursePlanId int

	if err := dbConn.QueryRow(context.Background(), query, req.StudentId).
		Scan(&startTime, &endTime, &coursePlanId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res.Error = "Student enrollment not found."
		}

		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	existenceCheckQuery := `
	SELECT true
	FROM attendance
	WHERE student_id = $1 AND course_plan_id = $2 AND date = CURRENT_DATE
	`
	doesExist := false

	if err := dbConn.QueryRow(context.Background(), existenceCheckQuery, req.StudentId, coursePlanId).Scan(&doesExist); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			res.Error = err.Error()
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
	}

	if doesExist == true {
		res.Error = "Attendance has already been logged."
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	now := time.Now()
	loc := now.Location()
	startTime = time.Date(now.Year(), now.Month(), now.Day(), startTime.Hour(), startTime.Minute(), startTime.Second(), 0, loc)
	endTime = time.Date(now.Year(), now.Month(), now.Day(), endTime.Hour(), endTime.Minute(), endTime.Second(), 0, loc)

	// TODO: Store this in a config file for configurability.
	var graceTime float64 = 50

	if (now.Equal(startTime) || now.After(startTime)) && (now.Equal(endTime) || now.Before(endTime)) {
		var err error

		if time.Since(startTime).Minutes() > graceTime {
			err = logLate(dbConn, req.StudentId, coursePlanId)

			if err != nil {
				res.Error = err.Error()
				c.AbortWithStatusJSON(http.StatusBadRequest, res)
				return
			}

			res.Message = "Attendance logged successfully."
			c.IndentedJSON(http.StatusOK, res)
			return
		}

		err = logPresent(dbConn, req.StudentId, coursePlanId)

		if err != nil {
			res.Error = err.Error()
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		res.Message = "Attendance logged successfully."
		c.IndentedJSON(http.StatusOK, res)
		return
	}

	res.Error = "Class session time has not entered. Attendance logging is disabled."
	c.AbortWithStatusJSON(http.StatusLocked, res)
}

func logPresent(dbConn *pgx.Conn, studentId int, coursePlanId int) error {
	command := `
	INSERT INTO attendance(student_id, course_plan_id, status_id)
	VALUES ($1, $2, $3)
	`
	// TODO: Find a way to load the distinct values of attendance statuses from the DB and keep them in memory through-out the application life. Use the in-memory attendance statuses to log attendance.

	// TODO: Should the context be the same as the one used in other places in this file?
	// TODO: Repalce 1 with a constant or enum.
	_, err := dbConn.Exec(context.Background(), command, studentId, coursePlanId, 1)

	if err != nil {
		return err
	}

	return nil
}

func logLate(dbConn *pgx.Conn, studentId int, coursePlanId int) error {
	command := `
	INSERT INTO attendance(student_id, course_plan_id, status_id)
	VALUES ($1, $2, $3)
	`
	_, err := dbConn.Exec(context.Background(), command, studentId, coursePlanId, 2)

	if err != nil {
		return err
	}

	return nil
}
