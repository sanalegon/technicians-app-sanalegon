package middlew

import (
	"net/http"

	"github.com/sanalegon/technicians-app-sanalegon/db"
)

/* CheckDB it the middleware that allows me to know the current state of the DB */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Connection with DB was lost", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
