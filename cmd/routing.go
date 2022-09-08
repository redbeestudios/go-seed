package cmd

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(dependencies *Dependencies) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pokemon/{name}", dependencies.PokemonController.GetPokemon).Methods(http.MethodGet)

	http.Handle("/", r)

	return r
}
