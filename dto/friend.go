package dto

type GetFriendResponse struct {
	FriendId        string  `json:"friendId"`
	PrivateRoomId   *string `json:"privateRoomId,omitempty"`
	LastReadStoryId int     `json:"lastReadStoryId"`
}

type GetFriendRequestResponse []string

type CreateFriendRequest struct {
	FriendId string `json:"friendId" binding:"required"`
}
