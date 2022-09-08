package cmd

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(dependencies *Dependencies) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pokemon", dependencies.PokemonController.GetPokemon)

	http.Handle("/", r)

	return r
}
