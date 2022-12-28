package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/out"
	"github.com/redbeestudios/go-seed/pkg"
)

var _ out.PokemonRepository = (*PokemonRestAdapter)(nil)

type PokemonRestAdapter struct {
	client *resty.Client
}

func (a *PokemonRestAdapter) SavePokemon(ctx context.Context, pokemon *pokemon.Pokemon) error {
	//TODO implement me
	panic("implement me")
}

func NewPokemonRestAdapter(
	config pkg.RestClientConfig,
) *PokemonRestAdapter {

	client := resty.New().
		SetBaseURL(config.BaseUrl).
		SetTimeout(time.Duration(config.TimeoutMilliseconds) * time.Millisecond).
		SetRetryCount(config.RetryCount).
		SetRetryWaitTime(time.Duration(config.RetryWaitMilliseconds) * time.Millisecond)

	return &PokemonRestAdapter{
		client: client,
	}
}

func (a *PokemonRestAdapter) GetByName(
	ctx context.Context,
	name string,
) (*pokemon.Pokemon, error) {
	response, err := a.client.
		R().
		SetPathParam("name", name).
		Get("/pokemon/{name}")

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("pokemon not found")
	}

	var responseObject *pokemonResponse = &pokemonResponse{}
	if err := json.Unmarshal(response.Body(), responseObject); err != nil {
		return nil, err
	}

	return responseObject.ToDomain()
}
