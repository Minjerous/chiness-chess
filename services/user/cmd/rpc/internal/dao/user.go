package dao

import (
	"chess-user/model"
	"gorm.io/gorm"
)

type Option func(*gorm.DB)

func UserID(userID int64) Option {
	return func(db *gorm.DB) {
		db.Where(&model.User{Uid: userID})
	}
}
func UserName(name string) Option {
	return func(db *gorm.DB) {
		db.Where(&model.User{UserName: name})
	}
}

func UserEmail(email string) Option {
	return func(db *gorm.DB) {
		db.Where(&model.User{Email: email})
	}
}

func GetUserInfo(options ...func(option *gorm.DB)) (infos model.User, err error) {
	DB = DB.Model(&model.User{})
	for _, option := range options {
		option(DB)
	}
	DB.Find(&infos)
	return infos, DB.Error
}

func InsertUser(user model.User) (err error) {
	return DB.Create(&user).Error
}
