package application

import (
	"your-project/api/entity"
	"your-project/api/repository"
	"your-project/api/response"
)

type UserUsecase interface {
	GetUser(userId string) (response.User, error)
	PostUser(user *entity.UserEntity) error
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{ur}
}

// [GET] /v1/user
func (uc *userUsecase) GetUser(userId string) (response.User, error) {

	user, err := uc.ur.GetUserInfo(userId)
	if err != nil {
		return response.User{}, err
	}

	return response.User{
		UserId:      user.UserId,
		MailAddress: user.MailAddress,
	}, nil
}

// [POST] /v1/user
func (uc *userUsecase) PostUser(user *entity.UserEntity) error {
	return nil
}
