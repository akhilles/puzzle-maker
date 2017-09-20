package main

import (
	"encoding/json"
	"fmt"
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
	rp, _, vm := puzzle.RandomPuzzle(n)
	fitness, minDist := puzzle.Evaluate(n, rp, vm)
	puzzle.PrintPuzzle(n, minDist)
	fmt.Println(fitness)

	json, _ := json.Marshal(rp)
	w.Write(json)
}
