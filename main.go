package main

import (
	"log"

	bd "github.com/sanalegon/technicians-app-sanalegon/db"
	"github.com/sanalegon/technicians-app-sanalegon/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("No connection with the database")
		return
	}
	handlers.Handler()
}
