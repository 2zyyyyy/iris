package services

import (
	"fmt"
	"iris/repositories"
)

type PokemonService interface {
	ShowPokemonName() string
}

type PokemonServiceManager struct {
	repo repositories.PokemonRepository
}

// ShowPokemonName 实现PokemonServiceManager的ShowPokemonName方法
func (p *PokemonServiceManager) ShowPokemonName() string {
	info := fmt.Sprintf("小精灵的名字是:%s", p.repo.GetPokemonName())
	fmt.Println(info)
	return info
}

func NewPokemonServiceManager(repo repositories.PokemonRepository) PokemonService {
	return &PokemonServiceManager{repo: repo}
}
