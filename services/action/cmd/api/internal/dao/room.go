package dao

import (
	"chess-game/model"
	"gorm.io/gorm"
)

type Option func(*gorm.DB)

func InsertRoom(room model.Room) error {
	return DB.Create(&room).Error
}

func RoomID(roomId int64) Option {
	return func(db *gorm.DB) {
		db.Where(&model.Room{Id: roomId})
	}
}

func GetRoomInfo(options ...func(option *gorm.DB)) (infos model.Room, err error) {
	DB = DB.Model(&model.Room{})
	for _, option := range options {
		option(DB)
	}
	DB.Find(&infos)
	return infos, DB.Error
}
