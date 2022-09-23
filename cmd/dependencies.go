package cmd

import (
	pokemonController "github.com/redbeestudios/go-seed/internal/adapter/controller/pokemon"
	"github.com/redbeestudios/go-seed/internal/adapter/redis"
	"github.com/redbeestudios/go-seed/internal/adapter/rest"
	useCase "github.com/redbeestudios/go-seed/internal/application/usecase/pokemon"
)

type Dependencies struct {
	PokemonController *pokemonController.PokemonController
}

func InitDependencies(config *Config) *Dependencies {

	pokemonRepository := rest.NewPokemonRestAdapter(config.PokeApi)
	cachedPokemonRepository := redis.NewCachedPokemonRestAdapter(
		config.Redis,
		pokemonRepository,
	)

	pokemonUseCase := useCase.NewGetByName(cachedPokemonRepository)

	pokemonController := pokemonController.NewPokemonController(pokemonUseCase)

	return &Dependencies{
		PokemonController: pokemonController,
	}
}
