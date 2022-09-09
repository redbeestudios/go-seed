package pokemon

type Pokemon struct {
	id   int
	name string
}

func NewPokemon(id int, name string) *Pokemon {
	return &Pokemon{id: id, name: name}
}

func (p *Pokemon) Id() int {
	return p.id
}

func (p *Pokemon) Name() string {
	return p.name
}
