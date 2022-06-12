package room

import (
	"action-game/cmd/api/internal/config"
	"action-game/cmd/api/internal/dao"
	"action-game/model"
	"chess-common/httpcode"
	"chess-common/jwt"
	"chess-common/rep"
	"chess-common/tool"
	"github.com/gin-gonic/gin"
	"strconv"
)

//创建房间
func Create(ctx *gin.Context) {

	token := ctx.PostForm("token")
	var Room model.Room
	err := ctx.ShouldBind(&Room)

	UserClaims, flag, _ := jwt.ParseToken(&jwt.UserClaims{
		AccessSecret:  config.GetUserCfg().Jwt.AccessSecret,
		RefreshSecret: config.GetUserCfg().Jwt.RefreshSecret,
	}, token)

	if err != nil {
		ctx.Set("Error", rep.Err{
			Code:   httpcode.StatusNotAcceptable,
			Msg:    "表单错误",
			Notice: nil,
		})
		ctx.Abort()

		return
	}

	if !flag {
		ctx.Set("Error", rep.Err{
			Code:   httpcode.StatusNotAcceptable,
			Msg:    "token 错误",
			Notice: nil,
		})
		ctx.Abort()

		return
	}

	err = dao.InsertRoom(model.Room{
		UserOne: UserClaims.Uid,
		RoomID:  Room.RoomID,
		Name:    Room.Name,
	})

	if err != nil {
		ctx.Set("Error", rep.Err{
			Code:   httpcode.StatusNotAcceptable,
			Msg:    "token error Or expired",
			Notice: nil,
		})
		ctx.Abort()

		return
	}

	var roomId int64
	dao.DB.Select("id").Find(&roomId)
	ctx.JSON(200, gin.H{
		"msg":       "房间创建成功",
		"code":      "200",
		"room_name": Room.Name,
		"room_id":   roomId,
	})

}

func EnterWS(ctx *gin.Context) {
	token := ctx.PostForm("token")
	roomId, err := strconv.ParseInt(ctx.Query("room_id"), 10, 10)

	if err != nil {
		tool.RespErrorWithData(ctx, "roomId 错误")

		return
	}

	UserClaims, flag, _ := jwt.ParseToken(&jwt.UserClaims{
		AccessSecret:  config.GetUserCfg().Jwt.AccessSecret,
		RefreshSecret: config.GetUserCfg().Jwt.RefreshSecret,
	}, token)

	if !flag {
		tool.RespErrorWithData(ctx, "token is error  OR expired")

		return
	}

	infos, err := dao.GetRoomInfo(dao.RoomID(roomId))

	if infos.UserOne == UserClaims.Uid {
		tool.RespSuccessfulWithData(ctx, "成功进入房间")

		return
	}

	if infos.UserTwo != 0 {
		tool.RespErrorWithData(ctx, "对局已经满人")

		return
	}

	if err != nil {
		tool.RespErrorWithData(ctx, "room 不存在")

		return
	}
	tool.RespSuccessfulWithData(ctx, "加入房间成功")

}
