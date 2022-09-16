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

var _ out.PokemonRepository = (*pokemonRestAdapter)(nil)

type typeDescription struct {
	Name string `json:"name"`
}

type responseType struct {
	Slot string          `json:"slot"`
	Type typeDescription `json:"type"`
}

type pokemonResponse struct {
	Id    int            `json:"id"`
	Name  string         `json:"name"`
	Types []responseType `json:"types"`
}

type pokemonRestAdapter struct {
	client *resty.Client
}

func NewPokemonRestAdapter(
	config pkg.RestClientConfig,
) *pokemonRestAdapter {

	client := resty.New().
		SetBaseURL(config.BaseUrl).
		SetTimeout(time.Duration(config.TimeoutMilliseconds) * time.Millisecond).
		SetRetryCount(config.RetryCount).
		SetRetryWaitTime(time.Duration(config.RetryWaitMilliseconds) * time.Millisecond)

	return &pokemonRestAdapter{
		client: client,
	}
}

func (a *pokemonRestAdapter) GetByName(
	ctx context.Context,
	name string,
) (*pokemon.Pokemon, error) {
	response, err := a.client.
		NewRequest().
		SetPathParam("name", name).
		Get("/pokemon/{name}")

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("pokemon not found")
	}

	var responseObject *pokemonResponse
	if err := json.Unmarshal(response.Body(), &responseObject); err != nil {
		return nil, fmt.Errorf("Error reading PokeAPI response")
	}

	return pokemon.NewPokemon(
		responseObject.Id,
		responseObject.Name,
		responseObject.Types[0].Type.Name,
	), nil
}
