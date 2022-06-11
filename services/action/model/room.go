package model

type Room struct {
	Id       int64  `gorm:"primarykey" "`
	Name     string `form:"room_name"`
	UserOne  int64
	UserTwo  int64
	PassWord string `form:"password"`
	RoomID   int64  `form:"room_id"`
}
