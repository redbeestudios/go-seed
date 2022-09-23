package pokemon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPokemon(t *testing.T) {
	id := 3
	name := "venusaur"
	pokemonType := MustBuildPokemonType(Grass.String())

	p := NewPokemon(id, name, pokemonType)

	assert.Equal(t, id, p.Id())
	assert.Equal(t, name, p.Name())
	assert.Equal(t, pokemonType, p.Type())
}
