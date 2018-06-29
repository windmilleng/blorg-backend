package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	r.HandleFunc("/pong", Pong)
	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "I'm the backend server.")
}

func Pong(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "pong")
}
