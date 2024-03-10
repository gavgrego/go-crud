package main

import (
	"api/routerFuncs"

	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// always close db connection after using DB
	defer db.Close()

	// create table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")

	if err != nil {
		log.Fatal(err)
	}

	// create router
	router := mux.NewRouter()
	router.HandleFunc("/users", routerFuncs.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", routerFuncs.GetUserByID(db)).Methods("GET")
	router.HandleFunc("/users", routerFuncs.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", routerFuncs.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", routerFuncs.DeleteUser(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", jsonContentTypeMiddleware(router)))

}

func jsonContentTypeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}
