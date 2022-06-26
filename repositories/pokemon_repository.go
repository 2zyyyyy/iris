package repositories

import "iris/models"

// PokemonRepository 定义宝可梦接口
type PokemonRepository interface {
	GetPokemonName() string
}

type PokemonManager struct {
}

// GetPokemonName PokemonManager实现了GetPokemonName 等于实现了PokemonRepository接口
func (p *PokemonManager) GetPokemonName() string {
	pokemon := models.Pokemon{Name: "迷你龙"}
	return pokemon.Name
}

func NewPokemonManager() PokemonRepository {
	// 因为PokemonManager实现了PokemonRepository接口的方法所以能返回 PokemonManager
	return &PokemonManager{}
}
