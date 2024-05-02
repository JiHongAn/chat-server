package model

import "time"

type Friend struct {
	UserId          string    `gorm:"column:userId;primaryKey;index:userId_isFriend;"`
	FriendId        string    `gorm:"column:friendId;primaryKey;"`
	IsFriend        bool      `gorm:"column:isFriend;index:userId_isFriend;"`
	LastReadStoryId int       `gorm:"column:lastReadStoryId;"`
	PrivateRoomId   string    `gorm:"column:privateRoomId;"`
	createdAt       time.Time `gorm:"column:createdAt;"`
}

func (Friend) TableName() string {
	return "friend"
}
