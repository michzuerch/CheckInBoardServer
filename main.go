package main

import (
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"github.com/michzuerch/CheckInBoardServer/util"
)

func main() {
	// util.TestSQLBoiler()

	// util.TestRest()

	// util.TestGin()

	util.CheckDatabaseAccess()

}
