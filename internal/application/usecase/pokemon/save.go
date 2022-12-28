package pokemon

import (
	"context"

	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/in"
	"github.com/redbeestudios/go-seed/internal/application/port/out"
)

var _ in.SavePokemon = (*SavePokemon)(nil)

type SavePokemon struct {
	pokemonRepository out.PokemonRepository
}

func (c *SavePokemon) Save(ctx context.Context, pokemon *pokemon.Pokemon) error {
	return c.pokemonRepository.SavePokemon(ctx, pokemon)
}

func NewSavePokemon(pokemonRepository out.PokemonRepository) *SavePokemon {
	return &SavePokemon{
		pokemonRepository: pokemonRepository,
	}
}
