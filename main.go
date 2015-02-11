package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Where is my Server %s?", req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
