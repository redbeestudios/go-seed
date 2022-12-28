package in

import (
	"context"

	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
)

//go:generate mockgen -source=./get_by_name.go -package=mocks -destination=../../../../mocks/get_pokemon_by_name.go

type SavePokemon interface {
	Save(ctx context.Context, pokemon *pokemon.Pokemon) error
}
