package testdata

import "github.com/redbeestudios/go-seed/internal/application/model/pokemon"

func Pokemon() *pokemon.Pokemon {
	return pokemon.NewPokemon(
		3,
		"venusaur",
		pokemon.MustBuildPokemonType(pokemon.Grass.String()),
	)
}
