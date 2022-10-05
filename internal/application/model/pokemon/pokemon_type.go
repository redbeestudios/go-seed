package pokemon

import "fmt"

type Type string

const (
	Fire  Type = "fire"
	Grass Type = "grass"
	Water Type = "water"
)

var allowedTypes = map[string]Type{
	Fire.String():  Fire,
	Grass.String(): Grass,
	Water.String(): Water,
}

func NewPokemonType(pokemonType string) (Type, error) {
	if t, ok := allowedTypes[pokemonType]; ok {
		return t, nil
	}
	return "", fmt.Errorf("Invalid pokemon type: %s", pokemonType)
}

func (t Type) String() string {
	return string(t)
}

// Testing purpose functions/methods
func MustBuildPokemonType(pokemonType string) Type {
	p, err := NewPokemonType(pokemonType)
	if err != nil {
		panic(err)
	}
	return p
}
