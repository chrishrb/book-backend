package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/aws_s3"
	_ "github.com/lib/pq"

	"github.com/chrishrb/bachelor-thesis/implementation/infrastructure/application/serverless/pkg/config"
)

var (
	dbName            = os.Getenv("POSTGRES_DB")
	migrationS3Bucket = os.Getenv("MIGRATION_S3_BUCKET")
	db                *sql.DB
)

type MigrateEvent struct {
	Command string `json:"command"`
}

func init() {
	var err error
	db, err = sql.Open("postgres", config.DBConnectionString)
	if err != nil {
		panic(err)
	}
}

func handler(_ context.Context, event MigrateEvent) (string, error) {
	cmd := event.Command

	// Connect to db and set migrate path
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return "", err
	}

	// Prepare migration
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("s3://%s", migrationS3Bucket), dbName, driver)
	if err != nil {
		return "", err
	}

	// Migrate UP or DOWN
	switch strings.ToLower(cmd) {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "drop":
		err = m.Drop()
	default:
		return "", errors.New("send event {\"command\": \"up\"} to migrate UP/DOWN/DROP")
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully migrated %s", strings.ToUpper(cmd)), nil
}

func main() {
	lambda.Start(handler)
}
