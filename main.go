package main // import "github.com/grange74/where-is-my-server"

import (
	"encoding/json"
	"fmt"
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
	fmt.Fprintf(writer, "Where is my Server %s?", req.URL.Path[1:])
}

func handleUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set(CONTENT_TYPE, JSON_CONTENT_TYPE)
	writer.WriteHeader(http.StatusOK)

	user := User{"Test User 1", "testuser1@gmail.com"}

	if err := json.NewEncoder(writer).Encode(user); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/user", handleUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
