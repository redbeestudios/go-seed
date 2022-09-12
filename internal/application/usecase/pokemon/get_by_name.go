package pokemon

import (
	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/in"
	"github.com/redbeestudios/go-seed/internal/application/port/out"
)

var _ in.GetPokemonByName = (*getByName)(nil)

type getByName struct {
	pokemonRepository out.PokemonRepository
}

func NewGetByName(pokemonRepository out.PokemonRepository) *getByName {
	return &getByName{
		pokemonRepository: pokemonRepository,
	}
}

func (c *getByName) Get(name string) (*pokemon.Pokemon, error) {
	return c.pokemonRepository.GetByName(name)
}
