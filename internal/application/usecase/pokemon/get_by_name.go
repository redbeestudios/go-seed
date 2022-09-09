package pokemon

import (
	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/in"
	"github.com/redbeestudios/go-seed/internal/application/port/out"
)

var _ in.GetPokemonByName = (*GetByName)(nil)

type GetByName struct {
	PokemonRepository out.PokemonRepository
}

func (c *GetByName) Get(name string) (*pokemon.Pokemon, error) {
	return c.PokemonRepository.GetByName(name)
}
