package out

import (
	"context"

	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
)

//go:generate mockgen -source=./repository.go -package=mocks -destination=../../../../mocks/pokemon_repository.go

type PokemonRepository interface {
	GetByName(ctx context.Context, name string) (*pokemon.Pokemon, error)
}
