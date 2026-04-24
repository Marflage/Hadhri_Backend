package handlers

import (
	"context"
	"database/sql"
	"errors"
	constants "hadhri/Constants"
	db "hadhri/Db"
	dtos "hadhri/Dtos"
	requests "hadhri/WebApi/Requests"
	responses "hadhri/WebApi/Responses"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *gin.Context) {
	var req requests.SignIn
	res := responses.ApiResponse[responses.SignIn]{}

	if err := c.ShouldBindJSON(&req); err != nil {
		// TODO: Log.
		// TODO: Create a middleware that parses the errors and returns an array of error messages for readability.
		// TODO: Find a way to set separate error messages for each violation.
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	var err error

	dbConn, err := db.InitDb()

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	ctx := context.Background()

	getStudentQuery := `
		SELECT id, email, password
		FROM students
		WHERE email = $1
	`

	var studentEmail string
	var studentPassword string
	var studentId int
	var adminPassword *string

	if err := dbConn.QueryRow(
		ctx,
		getStudentQuery,
		req.Email,
	).Scan(&studentId, &studentEmail, &studentPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if adminPassword, err = isAdmin(dbConn, ctx, req.Email); err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					res.Error = "Invalid credentials."
					c.AbortWithStatusJSON(http.StatusBadRequest, res)
					return
				}

				// TODO: Log the actual error.
				res.Error = err.Error()
				c.AbortWithStatusJSON(http.StatusInternalServerError, res)
				return
			}
		} else {
			// TODO: Log.
			res.Error = err.Error()
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}
	}

	var passwordHash string

	if studentPassword != "" {
		passwordHash = studentPassword
	} else {
		passwordHash = *adminPassword
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password)); err != nil {
		// TODO: Should the message be changed to "invalid password"?
		res.Error = "Invalid credentials."
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	token, err := generateToken(studentId)

	if err != nil {
		// TODO: Log.
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	// TODO: Store the student ID somewhere for the duration the backend is running.

	res.Data.StudentId = studentId
	res.Data.Token = token
	res.Message = "Signed in successfully."

	c.IndentedJSON(http.StatusOK, res)
}

func isAdmin(dbConn *pgx.Conn, ctx context.Context, requestedEmail string) (*string, error) {
	getAdminQuery := `
		SELECT password
		FROM admins
		WHERE email = $1
	`

	var adminPassword string

	if err := dbConn.QueryRow(ctx, getAdminQuery, requestedEmail).Scan(&adminPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// TODO: Log.
			return nil, err
		}

		// TODO: Log.
		return nil, err
	}

	return &adminPassword, nil
}

func generateToken(studentId int) (string, error) {
	claims := dtos.CustomClaims{
		StudentId: studentId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    constants.Issuer,
			Audience:  jwt.ClaimStrings{constants.Audience},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(constants.ExpirationMinutes) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(constants.Jwtkey)
}
