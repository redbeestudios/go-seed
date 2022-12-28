package pokemon

type Pokemon struct {
	id                   int
	name                 string
	pokemonType          Type
	pokemonSecondaryType Type
}

func NewPokemon(id int, name string, pokemonType Type, pokemonType2 Type) *Pokemon {
	return &Pokemon{id: id, name: name, pokemonType: pokemonType, pokemonSecondaryType: pokemonType2}
}

func (p *Pokemon) Type() Type {
	return p.pokemonType
}

func (p *Pokemon) Id() int {
	return p.id
}

func (p *Pokemon) Name() string {
	return p.name
}

func (p *Pokemon) SecondaryType() Type {
	return p.pokemonSecondaryType
}
