package util

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/michzuerch/CheckInBoardServer/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	_ "modernc.org/sqlite"
)

func TestSqlBoiler() {
	db, err := sql.Open("sqlite", "/home/michzuerch/Source/CheckInBoardServer/checkinboard-testing.db")

	dieIf(err)

	err = db.Ping()
	dieIf(err)

	fmt.Println("Database connected")

	u := &models.Person{Firstname: "django"}

	err = u.Insert(context.Background(), db, boil.Infer())

	dieIf(err)

	fmt.Println("Person id:", u.ID)
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
