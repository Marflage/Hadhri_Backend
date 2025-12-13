package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

const databaseUrlEnvVar string = "DATABASE_URL"

var Conn *pgx.Conn

func InitDb() error {
	// TODO: Experiment loading the connection string from a config file.
	connStr := os.Getenv(databaseUrlEnvVar)

	if connStr == "" {
		// return "Error obtaining DB connection string. " + databaseUrlEnvVar + " is not set."
	}

	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		return err
	}

	Conn = conn

	return nil
}
