package main

import (
	"log"
	"net/http"

	"github.com/shivamtiwari/routes"
)

func main() {

	router := routes.MovieRoutes()

	http.Handle("/api", router)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8081", router))
}
