package pokemon

type Type int

const (
	Fire Type = iota + 1
	Grass
	Water
	Normal
	Flying
	Fighting
	Poison
	Electric
	Ground
	Rock
	Psychic
	Ice
	Bug
	Ghost
	Steel
	Dragon
	Dark
	Fairy
	Invalid
)

func NewPokemonType(pokemonType string) Type {
	switch pokemonType {
	case "fire":
		return Fire
	case "water":
		return Water
	case "grass":
		return Grass
	case "normal":
		return Normal
	case "flying":
		return Flying
	case "fighting":
		return Fighting
	case "poison":
		return Poison
	case "electric":
		return Electric
	case "ground":
		return Ground
	case "rock":
		return Rock
	case "psychic":
		return Psychic
	case "ice":
		return Ice
	case "bug":
		return Bug
	case "ghost":
		return Ghost
	case "steel":
		return Steel
	case "dragon":
		return Dragon
	case "dark":
		return Dark
	case "fairy":
		return Fairy
	default:
		return Invalid
	}
}

func (t Type) String() string {
	switch t {
	case Fire:
		return "fire"
	case Water:
		return "water"
	case Grass:
		return "grass"
	case Normal:
		return "normal"
	case Flying:
		return "flying"
	case Fighting:
		return "fighting"
	case Poison:
		return "poison"
	case Electric:
		return "electric"
	case Ground:
		return "ground"
	case Rock:
		return "rock"
	case Psychic:
		return "psychic"
	case Ice:
		return "ice"
	case Bug:
		return "bug"
	case Ghost:
		return "ghost"
	case Steel:
		return "steel"
	case Dragon:
		return "dragon"
	case Dark:
		return "dark"
	case Fairy:
		return "fairy"
	default:
		return ""
	}
}

// Testing purpose functions/methods
func MustBuildPokemonType(pokemonType string) Type {
	p := NewPokemonType(pokemonType)
	return p
}
