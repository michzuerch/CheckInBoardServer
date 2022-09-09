package main

//go:generate sqlboiler sqlite3 --wipe -c sqlboiler-Sqlite3.toml

import (
	"fmt"
	"log"
	"net/http"
	"os"

	//"github.com/michzuerch/CheckInBoardServer/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/michzuerch/CheckInBoardServer/config"
	"github.com/michzuerch/CheckInBoardServer/hello"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//util.CheckDatabaseAccess()
	fmt.Println(hello.Greet())
	fmt.Printf("Test String len: %d\n", config.StringLength("Hello"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	fmt.Println("Shell:", os.Getenv("SHELL"))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	var httpPort string = os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
