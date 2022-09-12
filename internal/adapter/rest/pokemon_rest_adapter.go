package rest

import (
	"encoding/json"
	"fmt"
	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"io"
	"log"
	"net/http"
	"strconv"
)

const url string = "https://pokeapi.co/api/v2/"

type PokemonRestAdapter struct{}

func (p *PokemonRestAdapter) GetPokemonById(id int) (*pokemon.Pokemon, error) {
	resp, err := http.Get(url + "pokemon/" + strconv.Itoa(id))

	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)

	var pokemon pokemon.Pokemon
	err = json.Unmarshal(response, &pokemon)
	if err != nil {
		return nil, err
	}

	log.Println(pokemon)

	return &pokemon, nil
}
