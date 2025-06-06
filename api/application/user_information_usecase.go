package application

import (
	"your-project/api/repository"
	"your-project/api/response"
)

type UserInformationUsecase interface {
	GetUserInfo(userId string) (response.UserInformation, error)
}

type userInformationUsecase struct {
	ur repository.UserInformationRepository
}

func NewUserInformationUsecase(ur repository.UserInformationRepository) UserInformationUsecase {
	return userInformationUsecase{ur}
}

// GetMessage は指定された ID のメッセージを取得する
func (uc userInformationUsecase) GetUserInfo(userId string) (response.UserInformation, error) {

	user, err := uc.ur.GetUserInfo(userId)
	if err != nil {
		return response.UserInformation{}, err
	}

	return response.UserInformation{
		UserId:      user.UserId,
		MailAddress: user.MailAddress,
	}, nil
}
