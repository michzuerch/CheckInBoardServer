package util

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/michzuerch/CheckInBoardServer/models"
	"github.com/volatiletech/sqlboiler/v4/boil"

	// justify for linter??
	_ "modernc.org/sqlite"
)

func TestSQLBoiler() {
	envPostgresServer := os.Getenv("POSTGRES_SERVER")
	envPostgresUser := os.Getenv("POSTGRES_USER")
	envPostgresPassword := os.Getenv("POSTGRES_PASSWORD")
	envPostgresDB := os.Getenv("POSTGRES_DB")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", envPostgresServer, 5432, envPostgresUser, envPostgresPassword, envPostgresDB)

	fmt.Printf("Databasebase connection string from env: '%v' \n", psqlconn)
	db, err := sql.Open("postgres", psqlconn)

	dieIf(err)

	err = db.Ping()
	dieIf(err)

	fmt.Println("Database connected")

	// Create person
	p := &models.Person{
		Firstname: "Benny",
		Lastname:  "Hill",
	}

	err = p.Insert(context.Background(), db, boil.Infer())

	dieIf(err)

	fmt.Printf("Created person: %v \n", p)

	// Find person
	// person, err := models.FindPerson()

	u := &models.Person{Firstname: "django"}

	err = u.Insert(context.Background(), db, boil.Infer())

	dieIf(err)

	fmt.Println("Person id:", u.ID)

	got, err := models.Persons().One(context.Background(), db)

	dieIf(err)

	fmt.Printf("Got user id: %v, Firstname: %v\n", got.ID, got.Firstname)

	// foundAgain, err := models.Persons(qm.Where("firstname = ?", got.Name)).One()
	if err != nil {
		return
	}
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
