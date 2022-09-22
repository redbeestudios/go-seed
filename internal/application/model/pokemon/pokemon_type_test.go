package pokemon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPokemonType(t *testing.T) {

	t.Run("Return Type if name is valid", func(t *testing.T) {
		pokemonType, err := NewPokemonType(Fire.String())

		assert.NoError(t, err)
		assert.Equal(t, pokemonType, Fire)
	})

	t.Run("Return error if type name is invalid", func(t *testing.T) {
		_, err := NewPokemonType("???")

		assert.Error(t, err)
	})

}
