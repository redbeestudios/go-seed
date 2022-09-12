package pokemon

import (
	"github.com/redbeestudios/go-seed/internal/application/domain/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/out"
)

type GetById struct {
	pokemonRepository out.PokemonRepositoryOutputPort
}

func (g GetById) Execute(id int) (*pokemon.Pokemon, error) {
	return g.pokemonRepository.GetById(id)
}
