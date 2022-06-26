package repositories

import (
	"fmt"
	"iris/models"
)

// TrainerRepository 定义训练家接口
type TrainerRepository interface {
	CombatPokemon() string // 收服宝可梦方法
}

type TrainerManager struct {
}

// CombatPokemon TrainerManager 实现了CombatPokemon方法也就是实现了TrainerRepository接口
func (t *TrainerManager) CombatPokemon() string {
	trainer := models.Trainer{
		Name:   "小霞",
		Age:    12,
		Region: "丰缘地区(Hoenn Region)",
		PokemonName: models.Pokemon{
			Name: "宝石海星",
		},
		Skill: "高压水炮",
	}
	info := fmt.Sprintf("来自%s的训练家%s的%s使用了%s\n", trainer.Region,
		trainer.Name, trainer.PokemonName, trainer.Skill)
	return info
}

func NewTrainerManager() TrainerRepository {
	// 因为TrainerManager实现了TrainerRepository接口的方法所以能返回 TrainerManager
	return &TrainerManager{}
}
