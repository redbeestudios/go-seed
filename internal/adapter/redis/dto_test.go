package redis

import (
	"testing"

	"github.com/redbeestudios/go-seed/testdata"
	"github.com/stretchr/testify/assert"
)

func TestToDomain(t *testing.T) {
	pokemon := testdata.Pokemon()
	baseDTO := &pokemonDTO{
		Id:   pokemon.Id(),
		Name: pokemon.Name(),
		Type: pokemon.Type().String(),
	}

	type test struct {
		name      string
		dto       func() *pokemonDTO
		expectErr bool
	}

	tests := []test{
		{
			name: "Converts dto to domain",
			dto: func() *pokemonDTO {
				return baseDTO
			},
		},
		{
			name: "Error if pokemon type is invalid",
			dto: func() *pokemonDTO {
				response := baseDTO
				response.Type = "???"
				return response
			},
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := test.dto().ToDomain()
			if test.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, pokemon, res)
			}
		})
	}

}

func TestFromDomain(t *testing.T) {
	pokemon := testdata.Pokemon()

	dto := fromDomain(pokemon)

	assert.Equal(t, pokemon.Id(), dto.Id)
	assert.Equal(t, pokemon.Name(), dto.Name)
	assert.Equal(t, pokemon.Type().String(), dto.Type)
}
