package cmd

import (
	pokemonController "github.com/redbeestudios/go-seed/internal/adapter/controller/pokemon"
	"github.com/redbeestudios/go-seed/internal/adapter/rest"
	"github.com/redbeestudios/go-seed/internal/adapter/sql"
	useCase "github.com/redbeestudios/go-seed/internal/application/usecase/pokemon"
)

type Dependencies struct {
	PokemonController *pokemonController.PokemonController
}

func InitDependencies(config *Config) *Dependencies {

	pokemonRepository := rest.NewPokemonRestAdapter(config.PokeApi)
	/*cachedPokemonRepository := redis.NewCachedPokemonRestAdapter(
		config.Redis,
		pokemonRepository,
	)*/
	sqlPokemonRepository := sql.NewSqlPokemonRestAdapter(
		pokemonRepository,
	)

	pokemonUseCase := useCase.NewGetByName(pokemonRepository)

	savePokemonUseCase := useCase.NewSavePokemon(sqlPokemonRepository)

	controller := pokemonController.NewPokemonController(pokemonUseCase, savePokemonUseCase)

	return &Dependencies{
		PokemonController: controller,
	}
}
