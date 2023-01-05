package pokemon

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/template"
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

func (c *PokemonController) DumpPokemonsFromFile(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	go func() {
		start := time.Now()

		file, err := os.Open("/home/rodrigo/pokemonList/pokemonList.csv")
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		scanner := bufio.NewScanner(file)

		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			pokemonIndex, _ := strconv.Atoi(strings.Split(scanner.Text(), ",")[0])
			c.retrieveAndSavePokemon(response, ctx, pokemonIndex, true, false)
		}

		file.Close()

		log.Printf("Execution Finalized, elapsed time: %s", time.Since(start))
	}()

}

func (c *PokemonController) DumpPokemons(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	go func() {
		start := time.Now()
		for i := 1; i <= 906; i++ {
			c.retrieveAndSavePokemon(response, ctx, i, true, false)
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
			c.retrieveAndSavePokemon(response, ctx, index, false, false)
			times[index] = time.Since(start).Milliseconds()
		}()
	}

	wg.Wait()
	for _, eachTime := range times {
		totalTime = totalTime + eachTime
	}
	log.Printf("Execution Finalized, avg time: %d milliseconds", totalTime/905)

}

func (c *PokemonController) DumpPokemonsFromFileGoRoutine(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var times []int64

	var totalTime int64 = 0

	var wg sync.WaitGroup

	start := time.Now()

	file, err := os.Open("/home/rodrigo/pokemonList/pokemonList.csv")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		wg.Add(1)
		pokemonIndex, _ := strconv.Atoi(strings.Split(scanner.Text(), ",")[0])
		go func() {
			defer wg.Done()
			c.retrieveAndSavePokemon(response, ctx, pokemonIndex, false, false)
			times = append(times, time.Since(start).Milliseconds())
		}()

	}

	wg.Wait()

	err = file.Close()
	if err != nil {
		return
	}

	for _, eachTime := range times {
		totalTime = totalTime + eachTime
	}
	log.Printf("Execution Finalized, avg time: %d milliseconds", totalTime/int64(len(times)))

}

func (c *PokemonController) DumpPokemonsFromFilePDF(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	go func() {
		start := time.Now()

		file, err := os.Open("/home/rodrigo/pokemonList/pokemonList.csv")
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		scanner := bufio.NewScanner(file)

		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			pokemonIndex, _ := strconv.Atoi(strings.Split(scanner.Text(), ",")[0])
			c.retrieveAndSavePokemon(response, ctx, pokemonIndex, true, true)
		}

		file.Close()

		log.Printf("Execution Finalized, elapsed time: %s", time.Since(start))
	}()

}

func (c *PokemonController) DumpPokemonsPDF(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	go func() {
		start := time.Now()
		for i := 1; i <= 906; i++ {
			c.retrieveAndSavePokemon(response, ctx, i, true, true)
		}

		log.Printf("Execution Finalized, elapsed time: %s", time.Since(start))
	}()

}

func (c *PokemonController) DumpPokemonsWithGoRoutinesPDF(response http.ResponseWriter, request *http.Request) {
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
			c.retrieveAndSavePokemon(response, ctx, index, false, true)
			times[index] = time.Since(start).Milliseconds()
		}()
	}

	wg.Wait()
	for _, eachTime := range times {
		totalTime = totalTime + eachTime
	}
	log.Printf("Execution Finalized, avg time: %d milliseconds", totalTime/905)

}

func (c *PokemonController) DumpPokemonsFromFileGoRoutinePDF(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var times []int64

	var totalTime int64 = 0

	var wg sync.WaitGroup

	start := time.Now()

	file, err := os.Open("/home/rodrigo/pokemonList/pokemonList.csv")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		wg.Add(1)
		pokemonIndex, _ := strconv.Atoi(strings.Split(scanner.Text(), ",")[0])
		go func() {
			defer wg.Done()
			c.retrieveAndSavePokemon(response, ctx, pokemonIndex, false, true)
			times = append(times, time.Since(start).Milliseconds())
		}()

	}

	wg.Wait()

	err = file.Close()
	if err != nil {
		return
	}

	for _, eachTime := range times {
		totalTime = totalTime + eachTime
	}
	log.Printf("Execution Finalized, avg time: %d milliseconds", totalTime/int64(len(times)))

}

func (c *PokemonController) retrieveAndSavePokemon(response http.ResponseWriter, ctx context.Context, i int, logProcessingPokemon bool, savePokemonsAsPDF bool) {
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

	if savePokemonsAsPDF {

		// Client code
		pdfg, _ := wkhtmltopdf.NewPDFGenerator()
		htmlfile, err := template.ParseFiles("/home/rodrigo/pokemonTemplate/pokemonTemplate.html")

		if err != nil {
			log.Fatal(err)
		}

		var templateBytes bytes.Buffer

		pokemonType := pokemon.Type().String()

		if pokemon.SecondaryType().String() != "" {
			pokemonType = fmt.Sprintf("%s, %s", pokemonType, pokemon.SecondaryType().String())
		}

		pokemonOutput := pokemonResponse{
			Id:    pokemon.Id(),
			Name:  pokemon.Name(),
			Type:  pokemonType,
			Image: pokemon.Image(),
		}

		err = htmlfile.Execute(&templateBytes, pokemonOutput)
		if err != nil {
			return
		}

		pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(templateBytes.Bytes())))
		pdfg.Dpi.Set(600)

		err = pdfg.Create()
		if err != nil {
			log.Fatal(err)
		}

		err = pdfg.WriteFile("/home/rodrigo/pokemonPDF/" + strconv.Itoa(pokemon.Id()) + " - " + pokemon.Name())
		if err != nil {
			return
		}

	} else {
		f, err := os.Create("/home/rodrigo/pokemons/" + strconv.Itoa(pokemon.Id()) + " - " + pokemon.Name())

		if err != nil {
			log.Fatal("Prepare failed:", err.Error())
			return
		}
		defer f.Close()

		f.Write([]byte(pokemon.ToString()))

		f.Sync()
	}

	_ = c.savePokemon.Save(ctx, pokemon)
}
