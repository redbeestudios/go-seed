package pokemon

import (
	"strconv"

	"net/http"

	"github.com/redbeestudios/go-seed/pkg"
)

type PokemonController struct {
}

func NewPokemonController() *PokemonController {
	return &PokemonController{}
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
	pokemon := pokemons[id]
	response.Write([]byte(pokemon))
}
