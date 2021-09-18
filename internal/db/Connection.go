package db

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	User     string `envconfig:"POSTGRES_USER"`
	Name     string `envconfig:"POSTGRES_DB"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
	Host     string `envconfig:"POSTGRES_HOST"`
	Port     string `envconfig:"POSTGRES_PORT"`
}

var dbpool *pgxpool.Pool

func Init() {
	var err error
	c := dbConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
	dbpool, err = pgxpool.Connect(context.Background(), psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}

func Close() {
	dbpool.Close()
}

func getConnection() *pgxpool.Pool {
	return dbpool
}

func dbConfig() Config {
	var c Config
	err := envconfig.Process("spacelight", &c)
	if err != nil {
		log.Panic(err)
	}
	return c
}
