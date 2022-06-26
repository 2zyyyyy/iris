### Go Web Iris 极速入门
#### 快速开始
创建本地项目`iris`(具体项目名称自己定义),在对应的根目录下执行`go get`安装需要使用的第三方库
>> go get -u go get -u github.com/kataras/iris/v12@master
#### MVC项目结构
```
│  docker-compose.yml
│  Dockerfile
│  go.mod
│  go.sum
│  main.go
│  readme.md
│
├─models
│      pokemon.go
│      trainer.go
│
├─repositories
│      pokemon_repository.go
│      trainer_repository.go
│
├─services
│      pokemon_service.go
│      trainer_service.go
│
└─web
    ├─controller
    │      pokemon_controller.go
    │      trainer_contraller.go
    │
    └─views
        ├─pokemon
        │      index.html
        │
        └─trainer
                index.html
```