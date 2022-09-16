package util

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/michzuerch/CheckInBoardServer/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	_ "modernc.org/sqlite"
)

func TestSqlBoiler() {
	env_db_type := os.Getenv("DB_TYPE")
	env_db_connect_string := os.Getenv("DB_CONNECT_STRING")

	fmt.Printf("Databasebase Type and connect string from env: '%v', '%v' \n", env_db_type, env_db_connect_string)
	db, err := sql.Open(env_db_type, env_db_connect_string)

	dieIf(err)

	err = db.Ping()
	dieIf(err)

	fmt.Println("Database connected")

	//Create person
	p := &models.Person{
		Firstname: "Benny",
		Lastname:  "Hill",
	}

	err = p.Insert(context.Background(), db, boil.Infer())

	dieIf(err)

	fmt.Printf("Created person: %v \n", p)

	//Find person
	//person, err := models.FindPerson()

	u := &models.Person{Firstname: "django"}

	err = u.Insert(context.Background(), db, boil.Infer())

	dieIf(err)

	fmt.Println("Person id:", u.ID)

	got, err := models.Persons().One(context.Background(), db)

	dieIf(err)

	fmt.Printf("Got user id: %v, Firstname: %v\n", got.ID, got.Firstname)

	//foundAgain, err := models.Persons(qm.Where("firstname = ?", got.Name)).One()

}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
