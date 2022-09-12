package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/out"
)

func NewPokemonRestAdapter() *pokemonRestAdapter {
	return &pokemonRestAdapter{}
}

var _ out.PokemonRepository = (*pokemonRestAdapter)(nil)

type pokemonRestAdapter struct{}

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

func (c *pokemonRestAdapter) GetByName(name string) (*pokemon.Pokemon, error) {
	response, err := http.Get(fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%s", name))

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("pokemon not found")
	}

	var responseObject *pokemonResponse
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &responseObject)

	pokemon := pokemon.NewPokemon(
		responseObject.Id,
		responseObject.Name,
		responseObject.Types[0].Type.Name,
	)

	return pokemon, nil
}
