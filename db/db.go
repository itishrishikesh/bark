package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // needed for postgres driver
)

// BarkPostgresDb wraps the sqlx.DB in a custom struct to use it as a receiver for query functions
type BarkPostgresDb struct {
	Client *sqlx.DB
}

// ConnectToDatabase Connects to the Postgres DB
func ConnectToDatabase() (*BarkPostgresDb, error) {
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
		return &BarkPostgresDb{}, fmt.Errorf("error connecting to db: %w", err)
	}
	return &BarkPostgresDb{Client: dbConn}, nil
}

func OpenDatabase() (*BarkPostgresDb, error) {
	dbConn, err := sqlx.Open("postgres", "postgres://vaibhavkaushal:vaibhavkaushal@127.0.0.1:5432/bark?sslmode=disable")

	if err != nil {
		return &BarkPostgresDb{}, fmt.Errorf("E#1KDW57 - error connecting to db: %w", err)
	}

	return &BarkPostgresDb{Client: dbConn}, nil
}

func (d *BarkPostgresDb) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}