package pokemon

import "fmt"

type Type string

const (
	Fire     Type = "fire"
	Grass    Type = "grass"
	Water    Type = "water"
	Normal   Type = "normal"
	Flying   Type = "flying"
	Fighting Type = "fighting"
	Poison   Type = "poison"
	Electric Type = "electric"
	Ground   Type = "ground"
	Rock     Type = "rock"
	Psychic  Type = "psychic"
	Ice      Type = "ice"
	Bug      Type = "bug"
	Ghost    Type = "ghost"
	Steel    Type = "steel"
	Dragon   Type = "dragon"
	Dark     Type = "dark"
	Fairy    Type = "fairy"
	Invalid  Type = ""
)

var allowedTypes = map[string]Type{
	Fire.String():     Fire,
	Grass.String():    Grass,
	Water.String():    Water,
	Normal.String():   Normal,
	Flying.String():   Flying,
	Fighting.String(): Fighting,
	Poison.String():   Poison,
	Electric.String(): Electric,
	Ground.String():   Ground,
	Rock.String():     Rock,
	Psychic.String():  Psychic,
	Ice.String():      Ice,
	Bug.String():      Bug,
	Ghost.String():    Ghost,
	Steel.String():    Steel,
	Dragon.String():   Dragon,
	Dark.String():     Dark,
	Fairy.String():    Fairy,
	Invalid.String():  Invalid,
}

func NewPokemonType(pokemonType string) (*Type, error) {
	if t, ok := allowedTypes[pokemonType]; ok {
		return &t, nil
	}
	return nil, fmt.Errorf("Invalid pokemon type: %s", pokemonType)
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
	return *p
}
