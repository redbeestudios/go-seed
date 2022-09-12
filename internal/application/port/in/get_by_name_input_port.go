package in

import (
	"github.com/redbeestudios/go-seed/internal/application/domain/pokemon"
)

type GetByIdInputPort interface {
	Execute(id int) (*pokemon.Pokemon, error)
}
