package service

import (
	"your-project/api/entity"
	"your-project/api/repository"
)

type UserInformationService struct {
	repo repository.UserInformationRepositoryInt
}

// NewUserInformationService は新しいサービスを作成し、リポジトリを依存として注入します
func NewUserInformationService(repo repository.UserInformationRepositoryInt) *UserInformationService {
	return &UserInformationService{repo: repo}
}

// GetUserInfo は、指定された userId に基づいてユーザー情報を取得します
func (s *UserInformationService) GetUserInfo(userId string) (entity.UserInformation, error) {
	// リポジトリからユーザー情報を取得
	userInfo, err := s.repo.GetUserInfo(userId)
	if err != nil {
		return entity.UserInformation{}, err
	}
	return userInfo, nil
}
