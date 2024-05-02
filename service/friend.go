package service

import (
	"chat-server/dto"
	"chat-server/errors"
	"chat-server/model"
	"gorm.io/gorm"
)

type FriendService struct {
	DB *gorm.DB
}

func NewFriendService(db *gorm.DB) FriendService {
	return FriendService{DB: db}
}

// GetFriends 친구 목록 조회
func (service *FriendService) GetFriends(userId string) ([]*dto.GetFriendResponse, error) {
	var friends []*model.Friend

	err := service.DB.Table("friend").
		Select("friendId", "privateRoomId", "lastReadStoryId").
		Where("userId = ? AND isFriend = true", userId).
		Scan(&friends).Error

	if err != nil {
		return nil, errors.DBError
	}

	data := []*dto.GetFriendResponse{}

	for _, friend := range friends {
		var privateRoomIdPtr *string
		if friend.PrivateRoomId != "" {
			privateRoomIdPtr = &friend.PrivateRoomId
		} else {
			privateRoomIdPtr = nil
		}

		data = append(data, &dto.GetFriendResponse{
			FriendId:        friend.FriendId,
			PrivateRoomId:   privateRoomIdPtr,
			LastReadStoryId: friend.LastReadStoryId,
		})
	}
	return data, nil
}

// GetFriend 친구 정보 조회
func (service *FriendService) GetFriend(userId string, friendId string) (string, error) {
	var friend *model.Friend

	// 친구 요청 전송, 수신 여부 확인
	err := service.DB.
		Where("((userId = ? AND friendId = ?) OR (userId = ? AND friendId = ?))", userId, friendId, friendId, userId).
		Find(&friend).
		Error

	if err != nil {
		return "", errors.DBError
	}

	status := ""
	if friend.UserId == "" {
		status = "none"
	} else if friend.IsFriend {
		status = "friend"
	} else if friend.UserId == userId {
		status = "received"
	} else {
		status = "sent"
	}
	return status, nil
}

// GetFriendRequests 친구 스토리 조회
func (service *FriendService) GetFriendRequests(userId string) (dto.GetFriendRequestResponse, error) {
	var friendRequests []*model.Friend

	err := service.DB.Table("friend").
		Select("friendId").
		Where("userId = ? AND isFriend = false", userId).
		Scan(&friendRequests).Error

	if err != nil {
		return nil, errors.DBError
	}

	friendIds := []string{}
	for _, friend := range friendRequests {
		friendIds = append(friendIds, friend.FriendId)
	}
	return friendIds, nil
}

// CreateFriendRequest 친구 요청 생성
func (service *FriendService) CreateFriendRequest(userId string, friendId string) error {
	var count int64

	// 친구 여부 확인
	err := service.DB.Model(&model.Friend{}).
		Where("userId = ? AND friendId = ? AND isFriend = ?", userId, friendId, true).
		Count(&count).
		Error

	if err != nil {
		return errors.DBError
	} else if count == 1 {
		return errors.InvalidRequest
	}

	// 친구 요청 전송, 수신 여부 확인
	err = service.DB.Model(&model.Friend{}).
		Where("((userId = ? AND friendId = ?) OR (userId = ? AND friendId = ?)) AND isFriend = false", userId, friendId, friendId, userId).
		Count(&count).
		Error

	if err != nil {
		return errors.DBError
	} else if count == 1 {
		return errors.InvalidRequest
	}

	// 친구 요청 전송하기
	err = service.DB.Create(&model.Friend{
		UserId:   friendId,
		FriendId: userId,
		IsFriend: false,
	}).Error

	if err != nil {
		return errors.DBError
	}
	return nil
}

// FriendRequestAccept 친구 요청 수락
func (service *FriendService) FriendRequestAccept(userId string, friendId string) error {
	var count int64

	// 친구 요청 전송, 수신 여부 확인
	err := service.DB.Model(&model.Friend{}).
		Where("userId = ? AND friendId = ? AND isFriend = false", userId, friendId).
		Count(&count).
		Error

	if err != nil {
		return errors.DBError
	} else if count == 0 {
		return errors.InvalidRequest
	}

	// 친구 수락
	transaction := service.DB.Begin()
	transaction.Create(&model.Friend{
		UserId:   friendId,
		FriendId: userId,
		IsFriend: true,
	})
	transaction.Model(&model.Friend{}).
		Where("userId = ? AND friendId = ?", userId, friendId).
		Update("isFriend", true)

	err = transaction.Commit().Error
	if err != nil {
		transaction.Rollback()
		return errors.DBError
	}
	return nil
}

// DeleteFriend 친구 삭제
func (service *FriendService) DeleteFriend(userId string, friendId string) error {
	var count int64

	// 친구 여부 확인
	err := service.DB.Model(&model.Friend{}).
		Where("userId = ? AND friendId = ? AND isFriend = true", userId, friendId).
		Count(&count).
		Error

	if err != nil {
		return errors.DBError
	} else if count == 0 {
		return errors.InvalidRequest
	}

	// 친구 삭제
	transaction := service.DB.Begin()
	transaction.
		Where("userId = ? AND friendId = ?", userId, friendId).
		Delete(&model.Friend{})
	transaction.
		Where("userId = ? AND friendId = ?", friendId, userId).
		Delete(&model.Friend{})

	err = transaction.Commit().Error
	if err != nil {
		transaction.Rollback()
		return errors.DBError
	}
	return nil
}
