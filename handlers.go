package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandlePokemonById(rw http.ResponseWriter, r *http.Request, ps []Pokemon) {

	vars := mux.Vars(r)
	pokemonId, err := strconv.Atoi(vars["id"])

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Id passed is not a number"))
		return
	}

	for _, p := range ps {
		if p.Id == pokemonId {

			jsonResponse, err := json.Marshal(p)

			if err != nil {
				return
			}

			rw.Write(jsonResponse)
		}
	}

}
