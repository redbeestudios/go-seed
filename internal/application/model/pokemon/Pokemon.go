package pokemon

type Pokemon struct {
	id          int
	name        string
	pokemonType string
}

func NewPokemon(id int, name string, pokemonType string) *Pokemon {
	return &Pokemon{id: id, name: name, pokemonType: pokemonType}
}

func (p *Pokemon) PokemonType() string {
	return p.pokemonType
}

func (p *Pokemon) Id() int {
	return p.id
}

func (p *Pokemon) Name() string {
	return p.name
}
