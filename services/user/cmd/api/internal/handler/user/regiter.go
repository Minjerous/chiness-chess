package user

import (
	"chess-common/rep"
	"chess-common/rest/errdef"
	"chess-user/cmd/api/internal/handler/backend"
	"chess-user/cmd/rpc/proto/user"
	"chess-user/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var errInfo rep.Err
	var userRegister = model.ReqRegister{}

	userRegister.ReadParam(ctx)

	if err := userRegister.Check(); err != errdef.Nil {
		errInfo = rep.ErrInfo(err)
		fmt.Println(err)
		ctx.Set("Error", errInfo)
		ctx.Abort()
		return
	}

	registerResp, err := backend.GetUserClient().UserRegister(ctx, &user.UserRegisterRequest{
		Username: userRegister.UserName,
		Password: userRegister.Password,
		Email:    userRegister.Email,
	})

	if err != nil {
		ctx.JSON(500, "err")

		return
	}

	ctx.JSON(int(registerResp.StatusCode), registerResp)
	return
}
