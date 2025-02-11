package application

import (
	"your-project/api/response"
	"your-project/api/service"
)

type UserInformationUsecase interface {
	GetUserInfo(userId string) (response.UserInformation, error)
}

type userInformationUsecase struct {
	service service.UserInformationService
}

func NewUserInformationUsecase(su service.UserInformationService) UserInformationUsecase {
	return userInformationUsecase{su}
}

// GetMessage は指定された ID のメッセージを取得する
func (uc userInformationUsecase) GetUserInfo(userId string) (response.UserInformation, error) {

	user, err := uc.service.GetUserInfo(userId)
	if err != nil {
		return response.UserInformation{}, err
	}

	return response.UserInformation{
		UserId:      user.UserId,
		MailAddress: user.MailAddress,
	}, nil
}
