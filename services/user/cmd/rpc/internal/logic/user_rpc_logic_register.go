package logic

import (
	"chess-common/cryptx"
	"chess-common/httpcode"
	"chess-user/cmd/rpc/internal/dao"
	"chess-user/cmd/rpc/proto/user"
	"chess-user/model"
	"strconv"
	"time"
)

func Register(req *user.UserRegisterRequest) (res *user.UserRegisterResponse, err error) {

	if err != nil {
		res.StatusCode = 500
		return
	}
	salt := strconv.FormatInt(time.Now().Unix(), 10)

	userRegister := model.User{
		UserName:  req.Username,
		Salt:      salt,
		PassWord:  cryptx.EncryptSalt(salt, req.Password),
		Email:     req.Email,
		Signature: "这里什么都没有",
	}

	err = dao.InsertUser(userRegister)
	if err != nil {
		res.StatusCode = 500
		res.StatusMsg = "注册失败"

		return
	}

	res = &user.UserRegisterResponse{
		StatusCode: httpcode.StatusOK,
		StatusMsg:  "注册成功,请重新登录",
		Token:      "nil",
	}
	return
}
