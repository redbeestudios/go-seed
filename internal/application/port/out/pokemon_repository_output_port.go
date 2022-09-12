package out

import "github.com/redbeestudios/go-seed/internal/application/domain/pokemon"

type PokemonRepositoryOutputPort interface {
	GetById(id int) (*pokemon.Pokemon, error)
}
