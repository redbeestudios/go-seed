package pokemon

import (
	"encoding/json"
	"net/http"

	"github.com/redbeestudios/go-seed/internal/application/port/in"
	"github.com/redbeestudios/go-seed/pkg"
)

type PokemonController struct {
	getPokemonByName in.GetPokemonByName
}

func NewPokemonController(getPokemonByName in.GetPokemonByName) *PokemonController {
	return &PokemonController{
		getPokemonByName: getPokemonByName,
	}
}

func (c *PokemonController) GetPokemon(
	response http.ResponseWriter,
	request *http.Request,
) {
	ctx := request.Context()

	name, err := pkg.GetStringFromPath("name", request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	pokemon, err := c.getPokemonByName.Get(ctx, name)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Creo que esta rara esta validacion, si el pokemon es nulo deberiamos
	// haber propagado un error antes
	if pokemon == nil {
		http.Error(response, err.Error(), http.StatusNotFound)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(fromDomain(pokemon))

}
