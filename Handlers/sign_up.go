package handlers

import (
	"context"
	"database/sql"
	"errors"
	db "hadhri/Db"
	requests "hadhri/Requests"
	responses "hadhri/Responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var req requests.SignUp
	res := &responses.ApiResponse{}

	if err := c.ShouldBind(&req); err != nil {
		// TODO: Log error
		// TODO: Create a util to make the error messages more readable and return that.
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res) //gin.H{"error": err.Error()})
		return
	}

	dbConn, err := db.InitDb()

	if err != nil {
		res.Error = "Error connecting to the DB."
		c.JSON(http.StatusInternalServerError, res)
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
			// TODO: Create a struct for error.
			res.Error = "Course plan not found."
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		res.Error = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	insertStudentQuery := `
			INSERT INTO students(first_name, last_name, email, phone_number, password)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
		`

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	var studentId int

	if err := dbConn.QueryRow(
		ctx,
		insertStudentQuery,
		req.FirstName,
		req.LastName,
		req.Email,
		req.PhoneNumber,
		string(passwordHash),
	).Scan(&studentId); err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	insertEnrollmentQuery := `
		INSERT INTO student_enrollments(student_id, course_plan_id, semester)
		VALUES ($1, $2, $3)
	`

	commandTag, err := dbConn.Exec(
		ctx,
		insertEnrollmentQuery,
		studentId,
		coursePlanId,
		req.Semester,
	)

	if err != nil {
		res.Error = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	if commandTag.RowsAffected() == 1 {
		// TODO: Create a response type and return that.
		res.Message = "Signed up successfully."
		c.JSON(http.StatusOK, res)
		return
	}

	res.Error = "RowsAffected != 1"
	c.IndentedJSON(http.StatusOK, res)
}
