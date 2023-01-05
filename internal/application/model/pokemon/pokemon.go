package pokemon

import "fmt"

type Pokemon struct {
	id                   int
	name                 string
	pokemonType          Type
	pokemonSecondaryType Type
	image                string
}

func NewPokemon(id int, name string, pokemonType Type, pokemonType2 Type, image string) *Pokemon {
	return &Pokemon{id: id, name: name, pokemonType: pokemonType, pokemonSecondaryType: pokemonType2, image: image}
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

func (p *Pokemon) Image() string {
	return p.image
}

func (p *Pokemon) ToString() string {
	return fmt.Sprintf("Number: %d, Name: %s, Type1: %s, Type2: %s, Image: %s", p.id, p.name, p.pokemonType, p.pokemonSecondaryType, p.image)
}
