package pokemon

import "github.com/redbeestudios/go-seed/internal/application/model/pokemon"

type pokemonResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"types"`
	Image string `json:"image"`
}

func fromDomain(pokemon *pokemon.Pokemon) *pokemonResponse {
	types := pokemon.Type().String()
	if pokemon.SecondaryType().String() != "" {
		types = types + ", " + pokemon.SecondaryType().String()
	}
	return &pokemonResponse{
		Id:    pokemon.Id(),
		Name:  pokemon.Name(),
		Type:  types,
		Image: pokemon.Image(),
	}
}
