package pokemon

import pokemon_domain "github.com/redbeestudios/go-seed/src/application/model/pokemon"

type PokemonDTO struct {
	Id int `json:"id"`
}

func FromDomain(p *pokemon_domain.Pokemon) *PokemonDTO {
	return &PokemonDTO{
		Id: p.Id(),
	}
}
