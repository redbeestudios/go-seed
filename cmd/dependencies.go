package cmd

import (
	pokemonController "github.com/redbeestudios/go-seed/internal/adapter/controller/pokemon"
	"github.com/redbeestudios/go-seed/internal/adapter/rest"
	useCase "github.com/redbeestudios/go-seed/internal/application/usecase/pokemon"
)

type Dependencies struct {
	PokemonController *pokemonController.PokemonController
}

func InitDependencies() *Dependencies {

	pokemonRepository := rest.PokemonRestAdapter{}
	pokemonUseCase := useCase.GetByName{PokemonRepository: &pokemonRepository}
	pokemonController := pokemonController.NewPokemonController(&pokemonUseCase)

	return &Dependencies{
		PokemonController: pokemonController,
	}
}
