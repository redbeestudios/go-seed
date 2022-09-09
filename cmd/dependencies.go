package cmd

import "github.com/redbeestudios/go-seed/internal/adapter/controller/pokemon"

type Dependencies struct {
	PokemonController *pokemon.PokemonController
}

func InitDependencies() *Dependencies {
	return &Dependencies{
		PokemonController: pokemon.NewPokemonController(),
	}
}
