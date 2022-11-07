package util

import (
	"database/sql"
	"fmt"
	"os"

	// justify for linter ??
	_ "github.com/lib/pq"
)

func CheckDatabaseAccess() {
	envPostgresServer := os.Getenv("POSTGRES_SERVER")
	envPostgresUser := os.Getenv("POSTGRES_USER")
	envPostgresPassword := os.Getenv("POSTGRES_PASSWORD")
	envPostgresDB := os.Getenv("POSTGRES_DB")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", envPostgresServer, 5432, envPostgresUser, envPostgresPassword, envPostgresDB)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
