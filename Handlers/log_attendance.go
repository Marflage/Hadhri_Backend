package handlers

import (
	"context"
	"errors"
	db "hadhri/Db"
	responses "hadhri/WebApi/Responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func LogAttendance(c *gin.Context) {
	res := &responses.ApiResponse[any]{}

	// TODO: Create a const for the magic string.
	studentId := c.GetInt("studentId")

	dbConn, err := db.InitDb()

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	result, err := getClassSessionStartAndEndTimes(studentId, dbConn)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			res.Error = "Student enrollment not found."
		} else {
			res.Error = err.Error()
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	isAlreadyLogged, err := isAttendanceAlreadyLogged(studentId, dbConn)

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	if isAlreadyLogged {
		res.Error = "Attendance has already been logged."
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	now := time.Now()
	loc := now.Location()

	result.ClassSessionStartTime = time.Date(now.Year(), now.Month(), now.Day(), result.ClassSessionStartTime.Hour(), result.ClassSessionStartTime.Minute(), result.ClassSessionStartTime.Second(), 0, loc)

	result.ClassSessionEndTime = time.Date(now.Year(), now.Month(), now.Day(), result.ClassSessionEndTime.Hour(), result.ClassSessionEndTime.Minute(), result.ClassSessionEndTime.Second(), 0, loc)

	// TODO: Store this in a config file for configurability.
	var lateBuffer float64 = 50

	if (now.Equal(result.ClassSessionStartTime) || now.After(result.ClassSessionStartTime)) && (now.Equal(result.ClassSessionEndTime) || now.Before(result.ClassSessionEndTime)) {
		var err error

		// TODO: Use the Duration type for checking whether late should be logged or present.
		if time.Since(result.ClassSessionStartTime).Minutes() > lateBuffer {
			err = logLate(dbConn, studentId, result.CoursePlanId)

			if err != nil {
				res.Error = err.Error()
				c.AbortWithStatusJSON(http.StatusBadRequest, res)
				return
			}

			res.Message = "Attendance logged successfully."
			c.IndentedJSON(http.StatusOK, res)
			return
		}

		err = logPresent(dbConn, studentId, result.CoursePlanId)

		if err != nil {
			res.Error = err.Error()
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		res.Message = "Attendance logged successfully."
		c.JSON(http.StatusOK, res)
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

func getClassSessionStartAndEndTimes(studentId int, dbConn *pgx.Conn) (*queryResult, error) {
	query := `
	SELECT ses.start_time, ses.end_time, cp.id
	FROM student_enrollments se
	INNER JOIN course_plans cp
	ON cp.id = se.course_plan_id
	INNER JOIN class_sessions ses
	ON ses.id = cp.class_session_id
	WHERE se.student_id = $1
	`

	result := queryResult{}

	if err := dbConn.QueryRow(context.Background(), query, studentId).
		Scan(&result.ClassSessionStartTime, &result.ClassSessionEndTime, &result.CoursePlanId); err != nil {
		return nil, err
	}

	return &result, nil
}

// TODO: Move this in a file for reusability in the IsAttendanceLogged handler.
func isAttendanceAlreadyLogged(studentId int, dbConn *pgx.Conn) (bool, error) {
	query := `
	SELECT true
	FROM attendance
	WHERE student_id = $1 AND date = CURRENT_DATE
	`
	isLogged := false

	if err := dbConn.QueryRow(context.Background(), query, studentId).Scan(&isLogged); err != nil {
		return isLogged, err
	}

	return isLogged, nil
}

type queryResult struct {
	ClassSessionStartTime time.Time
	ClassSessionEndTime   time.Time
	CoursePlanId          int
}
