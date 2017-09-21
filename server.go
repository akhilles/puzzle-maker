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
	router.HandleFunc("/random", random).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("assets/"))))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func random(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))

	p := puzzle.GeneticPuzzle(n, 200, 5000, 6, 0.20, 0.7)

	json, _ := json.Marshal(p)
	w.Write(json)
}
