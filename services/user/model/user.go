package model

type User struct {
	Uid           int64 `gorm:"primarykey"`
	UserName      string
	Salt          string
	Phone         string
	Coin          int64
	PassWord      string
	Email         string
	FollowCount   int64
	FollowerCount int64
	Signature     string
	Avatar        string
}
