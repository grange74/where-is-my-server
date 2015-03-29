package main // import "github.com/grange74/where-is-my-server"

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONTENT_TYPE      = "Content-Type"
	JSON_CONTENT_TYPE = "application/json; charset=UTF-8"
)

type User struct {
	Name  string
	Email string
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Where is my Server?")
}

func UserShowHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set(CONTENT_TYPE, JSON_CONTENT_TYPE)
	rw.WriteHeader(http.StatusOK)

	username := mux.Vars(r)["username"]

	user := User{username, username + "@gmail.com"}

	if err := json.NewEncoder(rw).Encode(user); err != nil {
		panic(err)
	}
}

func main() {

	// had to use external Router to support parameters in path
	// chose Mux as it was one of the most use and most powerful.
	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/", HomeHandler)
	user := r.Path("/users/{id}").Subrouter()
	user.Methods("GET").HandlerFunc(UserShowHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
