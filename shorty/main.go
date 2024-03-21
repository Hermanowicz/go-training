package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"time"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Starting shorty app on port:", port)
}

func main() {
	// triger: true
	dbConn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err.Error())
	}
	dbConn.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = CreateTableAndIndex()
	if err != nil {
		log.Fatal("Database creation failed, terminating program")
	}

	// creaating new chi router
	r := chi.NewRouter()

	// Global middleware stack
	r.Use(middleware.CleanPath)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	s := http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    time.Second * 5,
		WriteTimeout:   time.Second * 5,
		IdleTimeout:    time.Second * 15,
		MaxHeaderBytes: 1 * 1024,
	}

	log.Fatal(s.ListenAndServe())
}
