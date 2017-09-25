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
	Cells      []int
	DepthBFS   []int
	Fitness    int
	Iterations int
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/genalgo", random).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("assets/"))))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func random(w http.ResponseWriter, r *http.Request) {
	// parameters
	// init pop = n * n * 2
	gens := 20000
	elitism := 8
	var selRate, mutRate float32
	selRate = 0.25
	mutRate = 0.018

	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	p, dbfs, fitness := puzzle.GeneticPuzzle(n, gens, elitism, selRate, mutRate)

	json, _ := json.Marshal(toGo{p, dbfs, fitness - n*n, gens})
	w.Write(json)
}
