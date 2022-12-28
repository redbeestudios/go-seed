package pokemon

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/redbeestudios/go-seed/internal/application/port/in"
	"github.com/redbeestudios/go-seed/pkg"
)

type PokemonController struct {
	getPokemonByName in.GetByName
	savePokemon      in.SavePokemon
}

func NewPokemonController(getPokemonByName in.GetByName, savePokemon in.SavePokemon) *PokemonController {
	return &PokemonController{
		getPokemonByName: getPokemonByName,
		savePokemon:      savePokemon,
	}
}

func (c *PokemonController) GetPokemon(
	response http.ResponseWriter,
	request *http.Request,
) {
	ctx := request.Context()

	name, err := pkg.GetStringFromPath("name", request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	pokemon, err := c.getPokemonByName.Get(ctx, name)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Creo que esta rara esta validacion, si el pokemon es nulo deberiamos
	// haber propagado un error antes
	if pokemon == nil {
		http.Error(response, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(fromDomain(pokemon))
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(js)
}

func (c *PokemonController) DumpPokemons(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	go func() {
		start := time.Now()
		for i := 1; i <= 906; i++ {
			c.retrieveAndSavePokemon(response, ctx, i, true)
		}

		log.Printf("Execution Finalized, elapsed time: %s", time.Since(start))
	}()

}

func (c *PokemonController) DumpPokemonsWithGoRoutines(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var times [906]int64

	var totalTime int64 = 0

	var wg sync.WaitGroup

	start := time.Now()
	for i := 1; i < 906; i++ {
		index := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.retrieveAndSavePokemon(response, ctx, index, false)
			times[index] = time.Since(start).Milliseconds()
		}()
	}

	wg.Wait()
	for _, eachTime := range times {
		totalTime = totalTime + eachTime
	}
	log.Printf("Execution Finalized, avg time: %d milliseconds", totalTime/905)

}

func (c *PokemonController) retrieveAndSavePokemon(response http.ResponseWriter, ctx context.Context, i int, logProcessingPokemon bool) {
	pokemon, err := c.getPokemonByName.Get(ctx, strconv.Itoa(i))
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	if logProcessingPokemon {
		log.Println("Processing: " + strconv.Itoa(pokemon.Id()) + " - " + pokemon.Name())
	}

	if pokemon == nil {
		http.Error(response, err.Error(), http.StatusNotFound)
		return
	}

	_ = c.savePokemon.Save(ctx, pokemon)
}
