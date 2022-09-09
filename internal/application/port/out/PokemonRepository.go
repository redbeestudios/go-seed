package out

import "github.com/redbeestudios/go-seed/internal/application/model/pokemon"

type PokemonRepository interface {
	GetByName(name string) (*pokemon.Pokemon, error)
}
