package services

import (
	"fmt"
	"iris/repositories"
)

// TrainerService 定义TrainerService解控
type TrainerService interface {
	ShowCombat() string
}

type TrainServiceManager struct {
	repo repositories.TrainerRepository
}

// ShowCombat 实现TrainServiceManager的ShowCombat方法
func (t *TrainServiceManager) ShowCombat() string {
	info := fmt.Sprintf("好的,现在战斗开始:%s", t.repo.CombatPokemon())
	fmt.Println(info)
	return info
}

func NewTrainServiceManager(repo repositories.TrainerRepository) TrainerService {
	return &TrainServiceManager{repo: repo}
}
