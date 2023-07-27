package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	return conn
}
