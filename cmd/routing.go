package cmd

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(dependencies *Dependencies) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pokemon/{name}", dependencies.PokemonController.GetPokemon).Methods(http.MethodGet)
	r.HandleFunc("/dumpPokemons", dependencies.PokemonController.DumpPokemons)
	r.HandleFunc("/dumpPokemonsGoRoutine", dependencies.PokemonController.DumpPokemonsWithGoRoutines)
	http.Handle("/", r)

	return r
}
