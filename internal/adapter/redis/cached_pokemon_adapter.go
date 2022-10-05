package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/out"
	"github.com/redbeestudios/go-seed/pkg"
)

var _ out.PokemonRepository = (*cachedPokemonRestAdapter)(nil)

const POKEMON_BY_NAME_KEY = "pokemon_with_name"

type cachedPokemonRestAdapter struct {
	cache      *redis.Client
	repository out.PokemonRepository
}

func NewCachedPokemonRestAdapter(
	cacheConfig pkg.RedisConfig,
	repository out.PokemonRepository,
) *cachedPokemonRestAdapter {
	cache := redis.NewClient(&redis.Options{
		Addr:     cacheConfig.BaseUrl,
		Password: "",
		DB:       0,
	})

	return &cachedPokemonRestAdapter{
		cache:      cache,
		repository: repository,
	}
}

func (a *cachedPokemonRestAdapter) GetByName(
	ctx context.Context,
	name string,
) (*pokemon.Pokemon, error) {
	var dto *pokemonDTO

	if cachedRes, err := a.cache.Get(ctx, keyForName(name)).Result(); err != redis.Nil {
		if err := json.Unmarshal([]byte(cachedRes), &dto); err != nil {
			return nil, fmt.Errorf(
				"Unmarshal error for pair (%s:%s), cause: %s",
				POKEMON_BY_NAME_KEY,
				name,
				err,
			)
		}

		return dto.ToDomain()
	}

	res, err := a.repository.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	dto = fromDomain(res)

	a.cache.Set(ctx, keyForName(name), dto, 0)

	return res, nil
}

func keyForName(name string) string {
	return POKEMON_BY_NAME_KEY + ":" + name
}
