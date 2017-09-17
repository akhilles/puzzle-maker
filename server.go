package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"puzzle-maker/puzzle"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/random", randomPuzzle).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("assets/"))))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func randomPuzzle(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	rp, _, _ := puzzle.RandomPuzzle(n)
	json, _ := json.Marshal(rp)
	w.Write(json)
}
