package handlers

import (
	"fmt"
	requests "hadhri/Requests"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	// TODO: Validate the incoming request data.

	var req requests.SignUpRequest

	if err := c.ShouldBind(&req); err != nil {
		// TODO: Log error
		// TODO: Create a util to make the error messages more readable and return that.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := fmt.Sprintln(req.FirstName, req.LastName, req.Email, req.PhoneNumber, req.CourseName, req.ClassSchedule, req.ClassSession, req.Password)
	response = strings.TrimSpace(response)

	{
		// TODO: Create a new student in the DB
		// dbConn, err := db.InitDb()

		// if err != nil {
		// 	// TODO: Return appropriate error.
		// 	// return err
		// 	fmt.Print("Error connecting to the db.")
		// }

		// dbConn.Exec(context.Background())
	}

	c.String(http.StatusOK, response)
}
