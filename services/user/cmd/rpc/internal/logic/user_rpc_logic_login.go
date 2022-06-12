package logic

import (
	"chess-common/cryptx"
	"chess-common/httpcode"
	"chess-common/jwt"
	"chess-user/cmd/rpc/internal/config"
	"chess-user/cmd/rpc/internal/dao"
	"chess-user/cmd/rpc/proto/user"
	"chess-user/model"
	"fmt"
)

func Login(req *user.UserLoginRequest) (res *user.UserLoginResponse, err error) {
	JwtCfg := config.GetUserCfg().Jwt
	var userModel model.User

	userModel, err = dao.GetUserInfo(dao.UserName(req.Account))

	if err != nil {
		res.StatusCode = 500
		return
	}
	req.Password = cryptx.EncryptSalt(userModel.Salt, req.Password)

	fmt.Println()
	if userModel.PassWord == req.Password {
		//r
		DoubleToken, err2 := jwt.GenToken(&jwt.UserClaims{
			Id:                userModel.Uid,
			AccessExpireTime:  68000,
			RefreshExpireTime: 680000,
			AccessSecret:      JwtCfg.AccessSecret,
			RefreshSecret:     JwtCfg.RefreshSecret,
		})

		if err2 != nil {
			res = &user.UserLoginResponse{
				StatusCode: httpcode.StatusInternalServerError,
				StatusMsg:  "StatusInternalServerError",
				UserId:     userModel.Uid,
				Token:      "nil",
			}
			return res, err
		}

		res = &user.UserLoginResponse{
			StatusCode: httpcode.StatusOK,
			StatusMsg:  "login success",
			UserId:     userModel.Uid,
			Token:      DoubleToken.AccessToken,
		}
		return
	} else {
		res = &user.UserLoginResponse{
			StatusCode: httpcode.StatusInternalServerError,
			StatusMsg:  "password is err",
			UserId:     userModel.Uid,
			Token:      "nil",
		}
		return
	}
	return
}
