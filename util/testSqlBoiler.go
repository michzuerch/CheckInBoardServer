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
	envDBType := os.Getenv("DB_TYPE")
	envDBConnection := os.Getenv("DB_CONNECTION")

	fmt.Printf("Databasebase Type and connect string from env: '%v', '%v' \n", envDBType, envDBConnection)
	db, err := sql.Open("postgres", envDBConnection)

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
