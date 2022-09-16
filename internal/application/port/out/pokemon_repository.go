package out

import (
	"context"

	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
)

type PokemonRepository interface {
	GetByName(ctx context.Context, name string) (*pokemon.Pokemon, error)
}
