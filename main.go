package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/michzuerch/CheckInBoardServer/hello"
	"github.com/michzuerch/CheckInBoardServer/util"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	fmt.Println("Test get shell from os.Getenv():", os.Getenv("SHELL"))
	util.TestSqlBoiler()
	fmt.Println(hello.Greet())
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

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
