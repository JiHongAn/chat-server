package service

import (
	"chat-server/dto"
	"chat-server/errors"
	"chat-server/model"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return UserService{DB: db}
}

func (service *UserService) CreateUser(dto *dto.CreateUser) error {
	user := model.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
	if result := service.DB.Create(&user); result.Error != nil {
		return errors.DBError
	}
	return nil
}

func (service *UserService) GetUser() ([]*dto.GetUserResponse, error) {
	var users []*model.User
	if result := service.DB.Find(&users); result.Error != nil {
		return nil, result.Error
	}

	var data []*dto.GetUserResponse
	for _, user := range users {
		data = append(data, &dto.GetUserResponse{
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return data, nil
}
