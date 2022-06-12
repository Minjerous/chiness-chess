package handler

import (
	"chess-game/cmd/api/internal/handler/room"
	"github.com/gin-gonic/gin"
)

type Room struct {
}

func (R *Room) Router(engine *gin.Engine) {
	UserGroup := engine.Group("chess/room")

	UserGroup.POST("/create", room.Create)

	UserGroup.POST("/enter", room.EnterWS)
	//UserGroup.POST("enter", room.Enter)

}
