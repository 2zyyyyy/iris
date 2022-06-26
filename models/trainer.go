package models

// Trainer 宝可梦训练家
type Trainer struct {
	Name        string
	Age         int
	Region      string
	PokemonName Pokemon
	Skill       string
}
