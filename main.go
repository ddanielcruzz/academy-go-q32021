package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Pokemon struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	TypeOne string `json:"typeOne"`
}

func main() {

	pokemons := initializePokemonData()

	r := mux.NewRouter()
	r.HandleFunc("/pokemon/{id}", func(rw http.ResponseWriter, r *http.Request) {
		HandlePokemonById(rw, r, pokemons)
	})
	http.ListenAndServe(":8080", r)
}

func initializePokemonData() []Pokemon {
	csvfile, err := os.Open("pokemons.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	csvR := csv.NewReader(csvfile)

	records, recordsErr := csvR.ReadAll()

	if recordsErr != nil {
		log.Fatalln("No records")
	}
	var pokemons []Pokemon
	for _, record := range records {
		idInt, _ := strconv.Atoi(record[0])
		data := Pokemon{
			Id:      idInt,
			Name:    record[1],
			TypeOne: record[2],
		}
		pokemons = append(pokemons, data)
	}
	return pokemons
}
