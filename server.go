package main // import "github.com/grange74/where-is-my-server"

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name  string
	Email string
}

func main() {

	// had to use external Router to support parameters in path
	// chose Mux as it was one of the most use and most powerful.
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/", HomeHandler)

	// Where is my server?
	where := router.Path("/where-is-my-server").Subrouter()
	where.Methods("GET").HandlerFunc(WhereIsMyServerHandler)

	// Server Resources
	servers := router.Path("/servers").Subrouter()
	servers.Methods("GET").HandlerFunc(ServersHandler)

	server := servers.Path("/servers/{id}").Subrouter()
	server.Methods("GET").HandlerFunc(GetServerHandler)
	server.Methods("POST").HandlerFunc(PostServerHandler)

	// Users Resource that belongs to a Server
	users := router.Path("/servers/{id}/users").Subrouter()
	users.Methods("GET").HandlerFunc(UsersHandler)
	user := router.Path("/servers/{id}/users/{id}").Subrouter()
	user.Methods("GET").HandlerFunc(GetUserHandler)
	user.Methods("POST").HandlerFunc(PostUserHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
