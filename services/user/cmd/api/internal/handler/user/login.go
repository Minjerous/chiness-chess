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

func Login(ctx *gin.Context) {
	var errInfo rep.Err
	var userLogin = model.ReqLogin{}

	userLogin.ReadParam(ctx)

	if err := userLogin.Check(); err != errdef.Nil {
		errInfo = rep.ErrInfo(err)
		fmt.Println(err)
		ctx.Set("Error", errInfo)
		ctx.Abort()
		return
	}

	fmt.Println(userLogin)

	loginResp, err := backend.GetUserClient().UserLogin(ctx, &user.UserLoginRequest{
		Account:  userLogin.Account,
		Password: userLogin.Password,
	})

	if err != nil {
		ctx.JSON(500, loginResp) //摆烂了
		fmt.Println(err)

		return
	}

	ctx.JSON(int(loginResp.StatusCode), loginResp)
	return

}
