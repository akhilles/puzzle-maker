package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"puzzle-maker/puzzle"

	"github.com/gorilla/mux"
)

type toGo struct {
	Cells    []int
	DepthBFS []int
	fitness  int
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/random", random).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("assets/"))))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func random(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	p := puzzle.GeneticPuzzle(n, 500, 40000, 20, 0.25, 4)
	fitness, dbfs := puzzle.Evaluate(n, p, true)

	json, _ := json.Marshal(toGo{p, dbfs, fitness})
	w.Write(json)
}
