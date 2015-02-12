package main // import "github.com/grange74/where-is-my-server"

import (
	"encoding/json"
	"fmt"
	"github.com/bmizerany/pat"
	"log"
	"net/http"
)

const (
	CONTENT_TYPE      = "Content-Type"
	JSON_CONTENT_TYPE = "application/json; charset=UTF-8"
)

type User struct {
	Name  string
	Email string
}

func handleRoot(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "Where is my Server?")
}

func handleUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set(CONTENT_TYPE, JSON_CONTENT_TYPE)
	writer.WriteHeader(http.StatusOK)

	username := req.URL.Query().Get(":username")

	user := User{username, username + "@gmail.com"}

	if err := json.NewEncoder(writer).Encode(user); err != nil {
		panic(err)
	}
}

func main() {

	// had to use external Router to support parameters in path
	// chose Pat as it was one of the simplest and fastest.
	m := pat.New()

	m.Get("/", http.HandlerFunc(handleRoot))
	m.Get("/users/:username", http.HandlerFunc(handleUser))

	http.Handle("/", m)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
