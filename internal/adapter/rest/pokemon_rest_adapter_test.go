package rest

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/redbeestudios/go-seed/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetByName(t *testing.T) {
	pokemon := testdata.Pokemon()

	type test struct {
		name      string
		mock      func()
		expectErr bool
	}

	tests := []test{
		{
			name: "Get pokemon by name",
			mock: func() {
				httpmock.RegisterResponder(
					"GET",
					fmt.Sprintf("/pokemon/%s", pokemon.Name()),
					httpmock.NewStringResponder(200, fmt.Sprintf(
						`{
							"id": %d,
							"name": "%s",
							"types": [{
								"type": {
									"name": "%s"
								}
							}]
						}`,
						pokemon.Id(),
						pokemon.Name(),
						pokemon.Type().String(),
					)))
			},
		},
		{
			name: "Error if client fails with 5xx",
			mock: func() {
				httpmock.RegisterResponder(
					"GET",
					fmt.Sprintf("/pokemon/%s", pokemon.Name()),
					httpmock.NewStringResponder(500, ""),
				)
			},
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := resty.New()
			if test.mock != nil {
				httpmock.ActivateNonDefault(client.GetClient())
				test.mock()
			}

			adapter := pokemonRestAdapter{client: client}

			res, err := adapter.GetByName(context.Background(), pokemon.Name())
			if test.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, pokemon, res)
			}
		})
	}

}
