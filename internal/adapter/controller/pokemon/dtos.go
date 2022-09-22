package pokemon

import "github.com/redbeestudios/go-seed/internal/application/model/pokemon"

type pokemonResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func fromDomain(pokemon *pokemon.Pokemon) *pokemonResponse {
	return &pokemonResponse{
		Id:   pokemon.Id(),
		Name: pokemon.Name(),
		Type: pokemon.Type().String(),
	}
}
