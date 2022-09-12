package in

import "github.com/redbeestudios/go-seed/internal/application/model/pokemon"

type GetPokemonByName interface {
	Get(name string) (*pokemon.Pokemon, error)
}
