package main

import (
	"chess-action/cmd/api/internal/config"
	"chess-action/cmd/api/internal/dao"
	"chess-action/cmd/api/internal/handler"
	"chess-common/rep"
	"chess-common/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	config.PareConfig()
	dao.InitDB()
	engine := gin.Default()
	engine.Use(rest.Cors(), rep.Recovery())
	routerEngine(engine)
	engine.Run(":8090")
}

func routerEngine(engine *gin.Engine) {
	new(handler.Room).Router(engine)
}
