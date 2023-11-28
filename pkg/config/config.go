package config

import (
	"fmt"
	"os"
)

var (
	dbHost     = os.Getenv("POSTGRES_HOST")
	dbPort     = os.Getenv("POSTGRES_PORT")
	dbUser     = os.Getenv("POSTGRES_USERNAME")
	dbPassword = os.Getenv("POSTGRES_PASSWORD")
	dbName     = os.Getenv("POSTGRES_DB")
	sslMode    = os.Getenv("SSL_MODE")

	DBConnectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode,
	)
)

const CorsAllowOrigin = "*"
