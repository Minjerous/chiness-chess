package model

import (
	"chess-common/httpcode/statuscode"
	"chess-common/rest/errdef"
	"chess-common/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReqLogin struct {
	Account  string
	Password string
}

func (r *ReqLogin) Check() errdef.Err {
	if len(r.Account) > 32 {
		return errdef.Errorf(http.StatusNotAcceptable, statuscode.CodeUnacceptedParam, "unexpect account, length should be less than 32")
	}
	if len(r.Password) > 32 {
		return errdef.Errorf(http.StatusNotAcceptable, statuscode.CodeUnacceptedParam, "unexpect password, length should be less than 32")
	}
	if len(r.Account)*len(r.Password) == 0 {
		return errdef.Errorf(http.StatusNotAcceptable, statuscode.CodeUnacceptedParam, "expect username and password")
	}
	return errdef.Nil
}

func (r *ReqLogin) ReadParam(ctx *gin.Context) {
	r.Account = tool.UnescapeString(ctx.PostForm("account"))
	r.Password = tool.UnescapeString(ctx.PostForm("password"))
}
