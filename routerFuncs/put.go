package routerFuncs

import (
	"api/data"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u data.User
		params := mux.Vars(r)
		id := params["id"]

		err := db.QueryRow("UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id", u.Name, u.Email, id).Scan(&u.ID)

		if err != nil {
			log.Fatal(err)
		}

		json.NewDecoder(r.Body).Decode(&u)
	}
}
