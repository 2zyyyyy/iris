package controller

import (
	"github.com/kataras/iris/v12/mvc"
	"iris/repositories"
	"iris/services"
)

type PokemonController struct {
}

func (p *PokemonController) Get() mvc.View {
	// 创建数据库操作对象
	pokemonRepository := repositories.NewPokemonManager()
	// 建立service
	pokemonService := services.NewPokemonServiceManager(pokemonRepository)
	// 根据service调用已实现的业务逻辑代码
	pokemonName := pokemonService.ShowPokemonName()
	return mvc.View{
		Name: "pokemon/index.html",
		Data: pokemonName,
	}
}
