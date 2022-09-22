package pokemon

type Pokemon struct {
	id          int
	name        string
	pokemonType Type
}

func NewPokemon(
	id int,
	name string,
	pokemonType Type,
) *Pokemon {
	return &Pokemon{
		id:          id,
		name:        name,
		pokemonType: pokemonType,
	}
}

func (p *Pokemon) Id() int {
	return p.id
}

func (p *Pokemon) Name() string {
	return p.name
}

func (p *Pokemon) Type() Type {
	return p.pokemonType
}
