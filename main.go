package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"github.com/michzuerch/CheckInBoardServer/hello"
	"github.com/michzuerch/CheckInBoardServer/util"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	fmt.Println("HOSTNAME os.Getenv(), HOSTNAME:", os.Getenv("HOSTNAME"))
	// util.TestSQLBoiler()

	// util.TestRest()

	// util.TestGin()

	util.CheckDatabaseAccess()

	fmt.Println("TestRest() done.")
	fmt.Println(hello.Greet())

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ping"))
	})

	var httpPort string = os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	http.ListenAndServe(httpPort, r)
}
