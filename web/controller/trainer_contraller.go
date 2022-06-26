package controller

import (
	"github.com/kataras/iris/v12/mvc"
	"iris/repositories"
	"iris/services"
)

type TrainerController struct {
}

func (p *TrainerController) Get() mvc.View {
	// 创建数据库操作对象
	trainerRepository := repositories.NewTrainerManager()
	// 建立service
	trainerService := services.NewTrainServiceManager(trainerRepository)
	// 根据service调用已实现的业务逻辑代码
	trainerInfo := trainerService.ShowCombat()
	return mvc.View{
		Name: "trainer/index.html",
		Data: trainerInfo,
	}
}
