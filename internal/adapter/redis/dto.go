package redis

import "github.com/redbeestudios/go-seed/internal/application/model/pokemon"

type pokemonDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (p *pokemonDTO) ToDomain() (*pokemon.Pokemon, error) {
	//pokemonType, err := pokemon.NewPokemonType(p.Type)
	//if err != nil {
	//	return nil, err
	//}

	return pokemon.NewPokemon(
		p.Id,
		p.Name,
		"water",
		"fire",
	), nil
}

func fromDomain(p *pokemon.Pokemon) *pokemonDTO {
	return &pokemonDTO{
		Id:   p.Id(),
		Name: p.Name(),
		Type: p.Type().String(),
	}
}
