package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONTENT_TYPE      = "Content-Type"
	JSON_CONTENT_TYPE = "application/json; charset=UTF-8"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Where is my Server?")
}

func WhereIsMyServerHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Where is my Server?")
}

func ServersHandler(rw http.ResponseWriter, r *http.Request) {
	servers := GetServersFromDB()
	if err := json.NewEncoder(rw).Encode(servers); err != nil {
		panic(err)
	}
}

func GetServerHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "This server")
}

func PostServerHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	ip := r.FormValue("ip")
	//TODO check that Name and IP are provided

	server := Server{name, ip}

	AddServerToDB(server)

	if err := json.NewEncoder(rw).Encode(server); err != nil {
		panic(err)
	}
}

func UsersHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "All the users")
}

func GetUserHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set(CONTENT_TYPE, JSON_CONTENT_TYPE)
	rw.WriteHeader(http.StatusOK)

	username := mux.Vars(r)["username"]

	user := User{username, username + "@gmail.com", "server1"}

	if err := json.NewEncoder(rw).Encode(user); err != nil {
		panic(err)
	}
}

func PostUserHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "New User created (NOT)!")
}
