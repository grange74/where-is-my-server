package main // import "github.com/grange74/where-is-my-server"

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	GET  = "GET"
	POST = "POST"
)

func main() {

	NewDB() //initialise the database connection

	// had to use external Router to support parameters in path
	// chose Mux as it was one of the most use and most powerful.
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/", HomeHandler)
	// Where is my server?
	router.HandleFunc("/where-is-my-server", WhereIsMyServerHandler).Methods(GET)

	// Server Resources
	router.HandleFunc("/servers", ServersHandler).Methods(GET)
	server := router.Path("/servers/{id}").Subrouter()
	server.Methods(GET).HandlerFunc(GetServerHandler)
	server.Methods(POST).HandlerFunc(PostServerHandler)

	// User Resource that belongs to a Server
	router.HandleFunc("/servers/{id}/users", UsersHandler).Methods(GET)
	user := router.Path("/servers/{id}/users/{id}").Subrouter()
	user.Methods(GET).HandlerFunc(GetUserHandler)
	user.Methods(POST).HandlerFunc(PostUserHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
