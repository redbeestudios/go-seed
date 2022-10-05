package pokemon

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/redbeestudios/go-seed/mocks"
	"github.com/redbeestudios/go-seed/testdata"
	"github.com/stretchr/testify/assert"
)

type getByNameDependencies struct {
	repository *mocks.MockPokemonRepository
}

func makeGetByNameDependencies(t *testing.T) *getByNameDependencies {
	return &getByNameDependencies{
		repository: mocks.NewMockPokemonRepository(gomock.NewController(t)),
	}
}

func TestNewGetByName(t *testing.T) {
	dependencies := makeGetByNameDependencies(t)

	usecase := NewGetByName(dependencies.repository)
	assert.NotNil(t, usecase)
}

func TestGetByName(t *testing.T) {
	pokemon := testdata.Pokemon()

	type test struct {
		name      string
		mock      func(deps *getByNameDependencies)
		expectErr bool
	}

	tests := []test{
		{
			name: "Returns expected pokemon if name exists",
			mock: func(deps *getByNameDependencies) {
				deps.repository.EXPECT().
					GetByName(gomock.Any(), pokemon.Name()).
					Return(pokemon, nil)
			},
			expectErr: false,
		},
		{
			name: "Fails if no pokemon exists with name",
			mock: func(deps *getByNameDependencies) {
				deps.repository.EXPECT().
					GetByName(gomock.Any(), pokemon.Name()).
					Return(nil, fmt.Errorf("no pokemon found"))
			},
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dependencies := makeGetByNameDependencies(t)
			if test.mock != nil {
				test.mock(dependencies)
			}

			usecase := NewGetByName(dependencies.repository)

			res, err := usecase.Get(context.Background(), pokemon.Name())
			if test.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, pokemon, res)
			}

		})
	}

}
