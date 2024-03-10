package routerFuncs

import (
	"api/data"

	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users")

		if err != nil {
			log.Fatal(err)
		}

		users := []data.User{}

		for rows.Next() {
			var u data.User

			if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
			fmt.Printf("user %v", u)

			err := rows.Err()
			if err != nil {
				log.Fatal(err)
			}
		}

		json.NewEncoder(w).Encode(users)
	}
}

func GetUserByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		var u data.User
		row := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)

		if err := row; err != nil {
			// fix later, don't want a fatal error if bad id is passed
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}
