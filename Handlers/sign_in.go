package handlers

import (
	"context"
	"database/sql"
	"errors"
	constants "hadhri/Constants"
	db "hadhri/Db"
	requests "hadhri/Requests"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *gin.Context) {
	var req requests.SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// TODO: Log.
		// TODO: Create a middleware that parses the errors and returns an array of error messages for readability.
		// TODO: Find a way to set separate error messages for each violation.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	var err error

	dbConn, err := db.InitDb()

	if err != nil {
		c.Errors.JSON()
		return
	}

	ctx := context.Background()

	getStudentQuery := `
		SELECT email, password
		FROM students
		WHERE email = $1
	`

	var studentEmail string
	var studentPassword string
	var adminPassword *string

	if err := dbConn.QueryRow(
		ctx,
		getStudentQuery,
		req.Email,
	).Scan(&studentEmail, &studentPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if adminPassword, err = isAdmin(dbConn, ctx, req.Email); err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials."})
					return
				}

				// TODO: Log the actual error.
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			// TODO: Log.
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials."})
		return
	}

	token, err := generateToken(req.Email)

	if err != nil {
		// TODO: Log.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: Store the student ID somewhere for the duration the backend is running.

	c.String(http.StatusOK, token)
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

func generateToken(requestedEmail string) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    constants.Issuer,
		Subject:   requestedEmail,
		Audience:  jwt.ClaimStrings{constants.Audience},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(constants.ExpirationMinutes) * time.Minute)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(constants.Jwtkey)
}
