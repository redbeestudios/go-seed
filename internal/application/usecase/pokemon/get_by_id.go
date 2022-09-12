package pokemon

import (
	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/out"
)

type GetById struct {
	pokemonRepository out.PokemonRepositoryOutputPort
}

func (g GetById) execute(id int) (*pokemon.Pokemon, error) {
	return g.pokemonRepository.GetPokemonById(id)
}
