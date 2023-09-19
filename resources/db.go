package resources

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// BarkPostgresDb wraps the sqlx.DB in a custom struct to use it as a receiver for query functions
type BarkPostgresDb struct {
	Client *sqlx.DB
}

var BarkDb *BarkPostgresDb

func InitDatabase() error {
	// Connect to Postgres DB instance
	var err error
	BarkDb, err = OpenDatabase()
	if err != nil {
		return fmt.Errorf("E#1KDZOZ - Opening database failed. Error: %v\n", err)
	}

	// Ping DB
	if err = BarkDb.Ping(context.Background()); err != nil {
		return fmt.Errorf("E#1KDZPY - Opening database failed. Error: %v\n", err)
	}

	fmt.Println("E#1KDZG7 - successfully connected to database")
	return nil
}

func OpenDatabase() (*BarkPostgresDb, error) {
	databaseURL := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	dbConn, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		return &BarkPostgresDb{}, fmt.Errorf("E#1KDW57 - error connecting to db: %w", err)
	}

	return &BarkPostgresDb{Client: dbConn}, nil
}

func (d *BarkPostgresDb) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
