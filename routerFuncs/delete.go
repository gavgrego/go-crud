package routerFuncs

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		_, err := db.Exec("DELETE FROM users WHERE id = $1", id)

		if err != nil {
			log.Fatal(err)
		}
	}
}
