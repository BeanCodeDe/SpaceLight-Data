package config

import "os"

var (
	//Service
	LogLevel = os.Getenv("LOG_LEVEL")

	//Auth
	AuthLoginUrl    = os.Getenv("AUTH_LOGIN_URL")
	ServiceId       = os.Getenv("SERVICE_ID")
	ServicePassword = os.Getenv("SERVICE_PASSWORD")
	AuthUserUrl     = os.Getenv("AUTH_USER_URL")

	//Database
	PostgresUser     = os.Getenv("POSTGRES_USER")
	PostgresDB       = os.Getenv("POSTGRES_DB")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	PostgresHost     = os.Getenv("POSTGRES_HOST")
	PostgresPort     = os.Getenv("POSTGRES_PORT")
)
