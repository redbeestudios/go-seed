package pokemon

import (
	"context"

	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/in"
	"github.com/redbeestudios/go-seed/internal/application/port/out"
)

var _ in.GetByName = (*GetByName)(nil)

type GetByName struct {
	pokemonRepository out.PokemonRepository
}

func NewGetByName(pokemonRepository out.PokemonRepository) *GetByName {
	return &GetByName{
		pokemonRepository: pokemonRepository,
	}
}

func (c *GetByName) Get(ctx context.Context, name string) (*pokemon.Pokemon, error) {
	return c.pokemonRepository.GetByName(ctx, name)
}
