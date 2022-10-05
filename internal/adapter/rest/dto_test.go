package rest

import (
	"testing"

	"github.com/redbeestudios/go-seed/testdata"
	"github.com/stretchr/testify/assert"
)

func TestToDomain(t *testing.T) {
	pokemon := testdata.Pokemon()
	baseDTO := &pokemonResponse{
		Id:   pokemon.Id(),
		Name: pokemon.Name(),
		Types: []responseType{
			{
				Type: typeDescription{
					Name: pokemon.Type().String(),
				},
			},
		},
	}

	type test struct {
		name      string
		dto       func() *pokemonResponse
		expectErr bool
	}

	tests := []test{
		{
			name: "Converts dto to domain",
			dto: func() *pokemonResponse {
				return baseDTO
			},
		},
		{
			name: "Error if pokemon type is invalid",
			dto: func() *pokemonResponse {
				response := baseDTO
				response.Types[0].Type.Name = "???"
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
