package pokemon

import (
	"encoding/json"
	modelPokemon "github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"net/http"
	"strconv"

	"github.com/redbeestudios/go-seed/pkg"
)

type PokemonController struct {
}

func NewPokemonController() *PokemonController {
	return &PokemonController{}
}

//TODO: import this struct from model
type Pokemon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (c *PokemonController) GetPokemon(
	response http.ResponseWriter,
	request *http.Request,
) {
	name, err := pkg.GetStringFromPath("name", request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	pokemons := map[int]string{
		1: "Pikachu",
		2: "Charmander",
		3: "Charizard",
		4: "Raichu",
	}
	id, err := strconv.Atoi(name)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	pokemonModel := Pokemon{}
	pokemon := pokemons[id]

	pokemonModel.Id = id
	pokemonModel.Name = pokemon
	pokemonModel.Type = modelPokemon.GetPokemonType(modelPokemon.Fire).Name()
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(pokemonModel)

}
