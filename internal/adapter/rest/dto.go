package rest

import "github.com/redbeestudios/go-seed/internal/application/model/pokemon"

type typeDescription struct {
	Name string `json:"name"`
}

type responseType struct {
	Type typeDescription `json:"type"`
}

type pokemonResponse struct {
	Id    int            `json:"id"`
	Name  string         `json:"name"`
	Types []responseType `json:"types"`
}

func (p *pokemonResponse) ToDomain() (*pokemon.Pokemon, error) {

	pokemonType, err := pokemon.NewPokemonType(p.Types[0].Type.Name)
	if err != nil {
		return nil, err
	}

	return pokemon.NewPokemon(
		p.Id,
		p.Name,
		pokemonType,
	), nil
}
