package rest

import (
	"encoding/json"
	"fmt"
	"github.com/redbeestudios/go-seed/internal/adapter/rest/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const url string = "https://pokeapi.co/api/v2/"

func GetPokemonById(id int) (*model.Pokemon, error) {
	resp, err := http.Get(url + "pokemon/" + strconv.Itoa(id))

	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)

	var pokemon model.Pokemon
	err = json.Unmarshal(response, &pokemon)
	if err != nil {
		return nil, err
	}

	log.Println(pokemon)

	return &pokemon, nil
}
