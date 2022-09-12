package in

import (
	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
)

type GetByNameInputPort interface {
	execute(id int) (*pokemon.Pokemon, error)
}
