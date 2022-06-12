package main

import (
	"action-game/cmd/api/internal/config"
	"action-game/cmd/api/internal/dao"
	"action-game/cmd/api/internal/handler"
	"chess-common/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	config.PareConfig()
	dao.InitDB()
	engine := gin.Default()
	engine.Use(rest.Cors())
	routerEngine(engine)
	engine.Run(":8090")
}

func routerEngine(engine *gin.Engine) {
	new(handler.Room).Router(engine)
}
