package handlers

import (
	"context"
	"database/sql"
	"errors"
	db "hadhri/Db"
	requests "hadhri/WebApi/Requests"
	responses "hadhri/WebApi/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var req requests.SignUp
	res := responses.ApiResponse[responses.SignIn]{}

	if err := c.ShouldBind(&req); err != nil {
		// TODO: Log error
		// TODO: Create a util to make the error messages more readable and return that.
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	dbConn, err := db.InitDb()

	if err != nil {
		// TODO: Log.
		res.Error = "Error connecting to the DB."
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	ctx := context.Background()

	// TODO: Is a transaction needed here and why?

	getCoursePlanIdQuery := `
			SELECT cp.id
			FROM course_plans cp
				JOIN available_semesters a
					ON cp.id = a.course_plan_id
			WHERE course_id = $1
			AND cp.class_schedule_id = $2
			AND cp.class_session_id = $3
			AND cp.is_active = TRUE
			AND a.semester = $4
		`

	var coursePlanId int

	if err := dbConn.QueryRow(
		ctx,
		getCoursePlanIdQuery,
		req.CourseId,
		req.ClassScheduleId,
		req.ClassSessionId,
		req.Semester,
	).Scan(&coursePlanId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// TODO: Log.
			res.Error = "Course plan not found."
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}

		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	insertStudentCommand := `
			INSERT INTO students(first_name, last_name, email, phone_number, password)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
		`

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		// TODO: Log.
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	var studentId int

	if err := dbConn.QueryRow(
		ctx,
		insertStudentCommand,
		req.FirstName,
		req.LastName,
		req.Email,
		req.PhoneNumber,
		string(passwordHash),
	).Scan(&studentId); err != nil {
		// TODO: Log.
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	insertEnrollmentCommand := `
		INSERT INTO student_enrollments(student_id, course_plan_id, semester)
		VALUES ($1, $2, $3)
	`

	commandTag, err := dbConn.Exec(
		ctx,
		insertEnrollmentCommand,
		studentId,
		coursePlanId,
		req.Semester,
	)

	if err != nil {
		// TODO: Log.
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	if commandTag.RowsAffected() <= 0 {
		// TODO: Log.
		res.Error = "RowsAffected != 1"
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	token, err := generateToken(studentId)

	if err != nil {
		// TODO: Log.
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res.Data.StudentId = studentId
	res.Data.Token = token
	res.Message = "Signed up successfully."

	c.IndentedJSON(http.StatusOK, res)
}
