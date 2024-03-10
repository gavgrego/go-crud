package main

import (
	"api/data"
	"api/routerFuncs"

	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// always close db connection after using DB
	defer db.Close()

	// create router
	router := mux.NewRouter()
	router.HandleFunc("/users", routerFuncs.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", getUserByID(db)).Methods("GET")
	router.HandleFunc("/users", createUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", jsonContentTypeMiddleware(router)))

}

func jsonContentTypeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func createUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u data.User

		err := db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewDecoder(r.Body).Decode(&u)

	}
}

func updateUser(db *sql.DB) http.HandlerFunc {
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

func deleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		_, err := db.Exec("DELETE FROM users WHERE id = $1", id)

		if err != nil {
			log.Fatal(err)
		}
	}
}
