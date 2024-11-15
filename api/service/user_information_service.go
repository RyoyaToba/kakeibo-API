package service

import (
	"your-project/api/entity"
	"your-project/api/repository"
)

type UserInformationService interface {
	GetUserInfo(userId string) (entity.UserInformation, error)
}

type userInformationService struct {
	ur repository.UserInformationRepository
}

// NewUserInformationService は新しいサービスを作成し、リポジトリを依存として注入します
func NewUserInformationService(ur repository.UserInformationRepository) UserInformationService {
	return userInformationService{ur}
}

// GetUserInfo は、指定された userId に基づいてユーザー情報を取得します
func (us userInformationService) GetUserInfo(userId string) (entity.UserInformation, error) {
	// リポジトリからユーザー情報を取得
	userInfo, err := us.ur.GetUserInfo(userId)
	if err != nil {
		return entity.UserInformation{}, err
	}
	return userInfo, nil
}
