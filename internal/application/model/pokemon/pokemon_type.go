package pokemon

const (
	Fire Type = iota
	Grass
	Water
)

type Type int

type PokemonType struct {
	name string
}

func (t *PokemonType) Name() string {
	return t.name
}

var typeMap = map[Type]*PokemonType{
	Fire:  {name: "Fire"},
	Grass: {name: "Grass"},
	Water: {name: "Water"},
}

func GetPokemonType(pokemonType Type) *PokemonType {
	return typeMap[pokemonType]
}
