package model

import "time"

type Friend struct {
	UserId          string    `gorm:"column:userId;primaryKey;index:userId_isFriend;"`
	FriendId        string    `gorm:"column:friendId;primaryKey;"`
	IsFriend        bool      `gorm:"column:isFriend;index:userId_isFriend;default:false;"`
	LastReadStoryId int       `gorm:"column:lastReadStoryId;default:0;"`
	PrivateRoomId   *string   `gorm:"column:privateRoomId;null;"`
	createdAt       time.Time `gorm:"column:createdAt;"`
}

func (Friend) TableName() string {
	return "friend"
}
