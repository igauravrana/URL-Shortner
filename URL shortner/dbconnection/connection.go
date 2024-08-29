package dbconnection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connect := "postgres://postgres:gaurav@localhost:5432/urlshortner?sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Established")
}
