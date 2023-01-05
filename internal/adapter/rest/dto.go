package rest

import (
	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
)

type typeDescription struct {
	Name string `json:"name"`
}

type responseType struct {
	Type typeDescription `json:"type"`
}

type artCategory struct {
	FrontDefault string `json:"front_default"`
}
type artCategories struct {
	OffArtwork artCategory `json:"official-artwork"`
}

type sprites struct {
	Other artCategories `json:"other"`
}

type pokemonResponse struct {
	Id    int            `json:"id"`
	Name  string         `json:"name"`
	Types []responseType `json:"types"`
	Image sprites        `json:"sprites"`
}

func (p *pokemonResponse) ToDomain() (*pokemon.Pokemon, error) {

	pokemonType1 := pokemon.NewPokemonType(p.Types[0].Type.Name)

	var pokemonType2 pokemon.Type
	if len(p.Types) > 1 {
		pokemonType2 = pokemon.NewPokemonType(p.Types[1].Type.Name)
	} else {
		pokemonType2 = pokemon.Invalid
	}

	return pokemon.NewPokemon(
		p.Id,
		p.Name,
		pokemonType1,
		pokemonType2,
		p.Image.Other.OffArtwork.FrontDefault,
	), nil
}
