package pokemon

type Pokemon struct {
	id int
}

func NewPokemon(id int) *Pokemon {
	return &Pokemon{id: id}
}

func (p *Pokemon) Id() int {
	return p.id
}
