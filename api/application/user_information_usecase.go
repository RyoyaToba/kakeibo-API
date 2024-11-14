package application

import (
	"database/sql"
	"your-project/api/entity"
	"your-project/api/service"
)

type UserInformationUsecase interface {
	GetUserInfo(userId string) (entity.UserInformation, error)
}

type UserInformationUsecaseImpl struct {
	service *service.UserInformationService
}

func NewUserInformationUsecase(db *sql.DB) *UserInformationUsecaseImpl {
	service := service.NewUserInformationService(db)
	return &UserInformationUsecaseImpl{service: service}
}

// GetMessage は指定された ID のメッセージを取得する
func (uc *UserInformationUsecaseImpl) GetUserInfo(userId string) (entity.UserInformation, error) {
	return uc.service.GetUserInfo(userId)
}
