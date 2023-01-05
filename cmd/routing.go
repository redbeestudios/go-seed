package cmd

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(dependencies *Dependencies) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pokemon/{name}", dependencies.PokemonController.GetPokemon).Methods(http.MethodGet)
	r.HandleFunc("/dump/basic/simple", dependencies.PokemonController.DumpPokemons)
	r.HandleFunc("/dump/basic/file", dependencies.PokemonController.DumpPokemonsFromFile)
	r.HandleFunc("/dump/basic/goRoutine", dependencies.PokemonController.DumpPokemonsWithGoRoutines)
	r.HandleFunc("/dump/basic/goRoutine/File", dependencies.PokemonController.DumpPokemonsFromFileGoRoutine)
	r.HandleFunc("/dump/pdf/simple", dependencies.PokemonController.DumpPokemonsPDF)
	r.HandleFunc("/dump/pdf/file", dependencies.PokemonController.DumpPokemonsFromFilePDF)
	r.HandleFunc("/dump/pdf/goRoutine", dependencies.PokemonController.DumpPokemonsWithGoRoutinesPDF)
	r.HandleFunc("/dump/pdf/goRoutine/File", dependencies.PokemonController.DumpPokemonsFromFileGoRoutinePDF)
	http.Handle("/", r)

	return r
}
