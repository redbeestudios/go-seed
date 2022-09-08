package pokemon

import (
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

	response.Write([]byte(name))
}
