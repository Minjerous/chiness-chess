package handler

import (
	"chess-user/cmd/api/internal/handler/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func (u *UserHandler) Router(engine *gin.Engine) {
	UserGroup := engine.Group("chess/user")

	UserGroup.POST("login", user.Login)
	UserGroup.POST("register", user.Register)
	UserGroup.Group("user-center", user.GetUserInfo)
	UserGroup.POST("user/login/auth", user.Register)
}
