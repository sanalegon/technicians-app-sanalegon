package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sanalegon/technicians-app-sanalegon/middlew"
	"github.com/sanalegon/twittor/routers"
	//"github.com/sanalegon/technicians/middlew"
	//"github.com/sanalegon/technicians/routers"
)

/* Handler sets current port, the handler and I tell the server to listen */
func Handler() {
	router := mux.NewRouter()

	// Routes(Endpoints):
	router.HandleFunc("/registro", middlew.CheckDB(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// use cors to make my API accessible everywhere
	handler := cors.AllowAll().Handler(router)
	// set port ant make it to start listening
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
