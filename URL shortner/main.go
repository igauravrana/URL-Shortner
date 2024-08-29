package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/igauravrana/URL-Shortner/dbconnection"
	"github.com/igauravrana/URL-Shortner/routes"
)

func main() {
	fmt.Println("Welcome to the URL shortner made by Gaurav")

	dbconnection.Connect()

	log.Fatal(http.ListenAndServe(":8080", routes.Router()))
}
