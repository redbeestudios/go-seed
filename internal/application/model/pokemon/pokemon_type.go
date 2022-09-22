package pokemon

import "fmt"

type Type string

const (
	Fire  Type = "Fire"
	Grass Type = "Grass"
	Water Type = "Water"
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
