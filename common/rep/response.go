package rep

import (
	"chess-common/rest/errdef"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Err struct {
	Code   int
	Msg    interface{}
	Notice interface{}
}
type Suc struct {
	Code   int
	Msg    interface{}
	Notice interface{}
}

func ErrInfo(errCode errdef.Err) Err {
	err := Err{
		errCode.HttpCode,
		errCode.Description,
		errCode.StatusCode,
	}
	return err
}

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err, flag := ctx.Get("Error") //得到返回的错误
		if flag {
			ctx.JSON(http.StatusOK, gin.H{
				"err": err.(Err),
			})
			ctx.Abort()
		}
		suc, flag := ctx.Get("Success")
		if flag {
			ctx.JSON(http.StatusOK, gin.H{
				"success": suc.(Suc),
			})
			ctx.Abort()
		}
		ctx.Abort()
		return
	}
}
