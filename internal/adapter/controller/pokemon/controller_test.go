package pokemon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/redbeestudios/go-seed/mocks"
	"github.com/redbeestudios/go-seed/testdata"
	"github.com/stretchr/testify/assert"
)

type controllerDependencies struct {
	getByName *mocks.MockGetByName
}

func makePokemonControllerDependencies(t *testing.T) *controllerDependencies {
	return &controllerDependencies{
		getByName: mocks.NewMockGetByName(gomock.NewController(t)),
	}
}

func TestNewPokemonController(t *testing.T) {
	dependencies := makePokemonControllerDependencies(t)

	controller := NewPokemonController(dependencies.getByName)
	assert.NotNil(t, controller)
}

func TestGetPokemon(t *testing.T) {
	pokemon := testdata.Pokemon()

	type test struct {
		name         string
		mock         func(*controllerDependencies)
		expectedBody string
		expectedCode int
	}

	tests := []test{
		{
			name: "prueba",
			mock: func(dependencies *controllerDependencies) {
				dependencies.getByName.EXPECT().
					Get(gomock.Any(), pokemon.Name()).
					Return(pokemon, nil)
			},
			expectedBody: `
				{
					"id": 3,
					"name": "venusaur",
					"type": "grass"
				}
			`,
			expectedCode: 200,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dependencies := makePokemonControllerDependencies(t)
			if test.mock != nil {
				test.mock(dependencies)
			}

			controller := NewPokemonController(dependencies.getByName)

			router := mux.NewRouter()
			router.HandleFunc("/pokemon/{name}", controller.GetPokemon).Methods(http.MethodGet)
			s := &http.Server{
				Handler: router,
			}

			req := httptest.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/pokemon/%s", pokemon.Name()),
				bytes.NewReader([]byte(test.expectedBody)),
			)
			rec := httptest.NewRecorder()
			s.Handler.ServeHTTP(rec, req)

			assert.Equal(t, test.expectedCode, rec.Code)

			if test.expectedBody != "" {
				buff := new(bytes.Buffer)
				err := json.Compact(buff, []byte(test.expectedBody))
				assert.NoError(t, err)
				assert.Equal(t, buff.String(), rec.Body.String())
			}

		})
	}

}
