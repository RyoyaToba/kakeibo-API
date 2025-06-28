package application

import (
	"errors"
	"your-project/api/entity"
	"your-project/api/repository"
	"your-project/api/request"
	"your-project/api/response"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	GetUser(userId string) (response.User, error)
	PostUser(ctx *gin.Context) (*response.User, error)
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
func (uc *userUsecase) PostUser(ctx *gin.Context) (*response.User, error) {

	// リクエストボディーの取得
	req := request.PostUserBodyRequest{}
	if err := req.Bind(ctx); err != nil {
		return nil, errors.New("faild request_body bind")
	}

	err := uc.ur.InsertUser(entity.UserEntity{
		UserId:      req.UserID,
		MailAddress: *req.MailAddress,
	})
	if err != nil {
		return nil, err
	}

	return &response.User{
		UserId:      req.UserID,
		MailAddress: *req.MailAddress,
	}, nil
}
