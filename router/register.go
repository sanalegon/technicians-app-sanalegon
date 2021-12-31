package routers

import (
	"encoding/json"
	"net/http"

	"github.com/sanalegon/technicians-app-sanalegon/db"
	"github.com/sanalegon/technicians-app-sanalegon/models"
)

// funciones devuelven nada, metodos si. Esto si es una funcion
/* Register is a function for creating the DB and the register for the users */
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t) // decodificamos lo que viene en el body dentro del modelo t. El body es un stream, es un dato que solo se puede leer una vez. Una vez leido, una vez destruido

	if err != nil {
		http.Error(w, "Error in the received data => "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email must be written", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password must contain at least 6 characters", 400)
		return
	}

	// la funcion devuelve tres valores, solo que ignorare el primero y el ultimo
	_, encontrado, _ := db.CheckIfUserExists(t.Email)

	if encontrado == true {
		http.Error(w, "There's already an User with the sent Email", 400)
		return
	}

	_, status, err := db.InsertRecord(t)

	if err != nil {
		http.Error(w, "It happened an error when trying to insert the current user => "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "It was not possible to insert the current user => ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
