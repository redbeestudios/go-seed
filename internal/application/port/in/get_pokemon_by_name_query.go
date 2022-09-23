package in

import (
	"context"

	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
)

type GetPokemonByName interface {
	Get(ctx context.Context, name string) (*pokemon.Pokemon, error)
}
