package database

import (
	"database/sql"
	"fmt"
	"go-template/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type Database struct {
	Client *sql.DB
}

func MakeDatabase(envs *config.Envs) *Database {

	USER := envs.DATABASE_USER
	PASSWORD := envs.DATABASE_PASSWORD
	HOST := envs.DATABASE_HOST
	PORT := envs.DATABASE_PORT
	NAME := envs.DATABASE_NAME

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER, PASSWORD, HOST, PORT, NAME))
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	return &Database{
		Client: db,
	}
}
