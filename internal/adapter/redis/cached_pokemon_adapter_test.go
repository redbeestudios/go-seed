package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/golang/mock/gomock"
	"github.com/redbeestudios/go-seed/mocks"
	"github.com/redbeestudios/go-seed/pkg"
	"github.com/redbeestudios/go-seed/testdata"
	"github.com/stretchr/testify/assert"
)

type cachedPokemonRestAdapterDependencies struct {
	redisConfig       pkg.RedisConfig
	pokemonRepository *mocks.MockPokemonRepository
}

func makeCachedPokemonRestAdapterDependencies(
	t *testing.T,
	miniredisClient *miniredis.Miniredis,
) *cachedPokemonRestAdapterDependencies {
	return &cachedPokemonRestAdapterDependencies{
		redisConfig: pkg.RedisConfig{
			BaseUrl: miniredisClient.Addr(),
		},
		pokemonRepository: mocks.NewMockPokemonRepository(gomock.NewController(t)),
	}
}

func TestNewCachedPokemonAdapter(t *testing.T) {
	miniredisClient := miniredis.RunT(t)
	deps := makeCachedPokemonRestAdapterDependencies(t, miniredisClient)

	adapter := NewCachedPokemonRestAdapter(deps.redisConfig, deps.pokemonRepository)

	assert.NotNil(t, adapter)
}

func TestGetByName(t *testing.T) {
	pokemon := testdata.Pokemon()

	type test struct {
		name           string
		mockRedis      func(*miniredis.Miniredis)
		mockRepository func(*mocks.MockPokemonRepository)
		expectErr      bool
	}

	tests := []test{
		{
			name: "Return cached pokemon if key exists in redis",
			mockRedis: func(client *miniredis.Miniredis) {
				client.Set(keyForName(pokemon.Name()), pkg.MustMarshal(fromDomain(pokemon)))
			},
			mockRepository: func(repository *mocks.MockPokemonRepository) {
				repository.EXPECT().
					GetByName(gomock.Any(), gomock.Any()).
					Times(0)
			},
		},
		{
			name: "Error if repository fails to return pokemon",
			mockRepository: func(repository *mocks.MockPokemonRepository) {
				repository.EXPECT().
					GetByName(gomock.Any(), pokemon.Name()).
					Return(nil, fmt.Errorf("No pokemon found"))
			},
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			miniredisClient := miniredis.RunT(t)
			deps := makeCachedPokemonRestAdapterDependencies(t, miniredisClient)
			adapter := NewCachedPokemonRestAdapter(deps.redisConfig, deps.pokemonRepository)

			if test.mockRedis != nil {
				test.mockRedis(miniredisClient)
			}
			if test.mockRepository != nil {
				test.mockRepository(deps.pokemonRepository)
			}

			res, err := adapter.GetByName(context.Background(), pokemon.Name())
			if test.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, pokemon, res)
			}
		})
	}

}
