package model

type Pokemon struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

func (p *Pokemon) Id() int {
	return p.id
}

func (p *Pokemon) Name() string {
	return p.name

}
