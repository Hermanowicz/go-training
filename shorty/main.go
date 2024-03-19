package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

func main() {
	dbConn, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal(err.Error())
	}

	dbConn.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(`Hello, world!
On multiple lines.
I hope you will have a great DAY.`)
	fmt.Println(time.Now().UTC())
}
