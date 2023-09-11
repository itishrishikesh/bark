package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //needed for postgres driver
)

// Wrapping the sqlx.DB in a custom struct to use it as a receiver for query functions in barkLogQueries.go
type Database struct {
	Client *sqlx.DB
}

// Connects to the Postgres DB
func ConnectToDatabase() (*Database, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL_MODE"),
	)

	dbConn, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		return &Database{}, fmt.Errorf("error connecting to db: %w", err)
	}
	return &Database{Client: dbConn}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
