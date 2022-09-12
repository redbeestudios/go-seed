package pokemon

import (
	"encoding/json"
	modelPokemon "github.com/redbeestudios/go-seed/internal/application/domain/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/in"
	"net/http"
	"strconv"

	"github.com/redbeestudios/go-seed/pkg"
)

type PokemonController struct {
	getById in.GetByIdInputPort
}

func NewPokemonController(getById in.GetByIdInputPort) *PokemonController {
	return &PokemonController{
		getById: getById,
	}
}

//TODO: import this struct from domain
type PokemonResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (c *PokemonController) GetById(
	response http.ResponseWriter,
	request *http.Request,
) {

	stringId, _ := pkg.GetStringFromPath("id", request)

	id, err := strconv.Atoi(stringId)

	pokemon, err := c.getById.Execute(id)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	if pokemon == nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	pokemonResponse := PokemonResponse{
		Id:   pokemon.Id(),
		Name: pokemon.Name(),
		Type: pokemon.PokemonType(),
	}

	json.NewEncoder(response).Encode(pokemonResponse)

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

	pokemonModel := PokemonResponse{}
	pokemon := pokemons[id]

	pokemonModel.Id = id
	pokemonModel.Name = pokemon
	pokemonModel.Type = modelPokemon.GetPokemonType(modelPokemon.Fire).Name()
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(pokemonModel)

}
