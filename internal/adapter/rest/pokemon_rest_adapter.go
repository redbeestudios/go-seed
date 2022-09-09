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

var _ out.PokemonRepository = (*PokemonRestAdapter)(nil)

type PokemonRestAdapter struct{}

type TypeDescription struct {
	Name string `json:"name"`
}

type ResponseType struct {
	Slot string          `json:"slot"`
	Type TypeDescription `json:"type"`
}

type PokemonResponse struct {
	Id    int            `json:"id"`
	Name  string         `json:"name"`
	Types []ResponseType `json:"types"`
}

func (c *PokemonRestAdapter) GetByName(name string) (*pokemon.Pokemon, error) {
	//var pokemonResponse PokemonResponse
	response, err := http.Get(fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%s", name))

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("pokemon not found")
	}

	var responseObject *PokemonResponse
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
