package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

const gamesPath = "/arachne/games/"

func newGameHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	r := mux.NewRouter()

	// REST: Create a new entry in the collection. The new entry's URI is assigned
	// automatically and is usually returned by the operation.
	r.HandleFunc(gamesPath, newGameHandler).Methods("POST")

	http.ListenAndServe(":8000", r)
}
