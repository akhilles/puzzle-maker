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
	Solution   []string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/genalgo", random).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("assets/"))))

	//puzzle.GetPlot(5, 10000, 0.3, 0.018, 10)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func random(w http.ResponseWriter, r *http.Request) {
	// parameters
	// init pop = n * n * 2
	gens := 160000
	var selRate, mutRate float32
	selRate = 0.3
	mutRate = 0.018

	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	p, dbfs, fitness, solution := puzzle.GeneticPuzzle(n, gens, selRate, mutRate)

	json, _ := json.Marshal(toGo{p, dbfs, fitness - n*n, gens, solution})
	w.Write(json)
}
