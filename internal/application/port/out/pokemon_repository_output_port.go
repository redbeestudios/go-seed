package out

import "github.com/redbeestudios/go-seed/internal/application/model/pokemon"

type PokemonRepositoryOutputPort interface {
	GetPokemonById(id int) (*pokemon.Pokemon, error)
}
