package model

import (
	"chess-common/httpcode/statuscode"
	"chess-common/rest/errdef"
	"chess-common/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReqRegister struct {
	Email    string
	UserName string
	Password string
}

func (r *ReqRegister) Check() errdef.Err {
	if !tool.VerifyEmailFormat(r.Email) {
		return errdef.Errorf(http.StatusNotAcceptable, statuscode.CodeUnacceptedParam, "email format is error")
	}
	if len(r.Password) > 32 || len(r.Password) < 6 {
		return errdef.Errorf(http.StatusNotAcceptable, statuscode.CodeUnacceptedParam, "unexpect password, length should be in 6~32")
	}
	if len(r.UserName)*len(r.Password) == 0 {
		return errdef.Errorf(http.StatusNotAcceptable, statuscode.CodeUnacceptedParam, "expect username and password")
	}
	return errdef.Nil
}

func (r *ReqRegister) ReadParam(ctx *gin.Context) {
	r.Email = tool.UnescapeString(ctx.PostForm("email"))
	r.Password = tool.UnescapeString(ctx.PostForm("password"))
	r.UserName = tool.UnescapeString(ctx.PostForm("username"))
}
