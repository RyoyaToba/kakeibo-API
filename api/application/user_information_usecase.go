package application

import (
	"your-project/api/entity"
	"your-project/api/service"
)

type UserInformationUsecase interface {
	GetUserInfo(userId string) (entity.UserInformation, error)
}

type userInformationUsecase struct {
	service service.UserInformationService
}

func NewUserInformationUsecase(su service.UserInformationService) UserInformationUsecase {
	return userInformationUsecase{su}
}

// GetMessage は指定された ID のメッセージを取得する
func (uc userInformationUsecase) GetUserInfo(userId string) (entity.UserInformation, error) {
	return uc.service.GetUserInfo(userId)
}
