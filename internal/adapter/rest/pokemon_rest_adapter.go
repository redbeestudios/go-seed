package rest

import (
	"encoding/json"
	"github.com/redbeestudios/go-seed/internal/adapter/rest/model"
	"io"
	"log"
	"net/http"
	"strconv"
)

const url = "https://pokeapi.co/api/v2/"

func GetPokemonById(id int) (*model.Pokemon, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url+"pokemon"+strconv.Itoa(id), nil)

	if err != nil {
		log.Println("No existe")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("No existe")
	}

	var pokemon model.Pokemon
	err = json.Unmarshal(bodyBytes, &pokemon)

	if err != nil {
		log.Println("No existe")
	}

	defer resp.Body.Close()

	return &pokemon, nil
}
