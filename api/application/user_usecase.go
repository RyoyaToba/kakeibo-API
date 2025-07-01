package application

import (
	"errors"
	"your-project/api/entity"
	"your-project/api/repository"
	"your-project/api/request"
	"your-project/api/response"
	"your-project/infra"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UserUsecase interface {
	GetUser(userId string) (response.User, error)
	PostUser(ctx *gin.Context) (*response.UserResponse, error)
}

type userUsecase struct {
	ur repository.UserRepository
	lr repository.LoginRepository
}

func NewUserUsecase(ur repository.UserRepository, lr repository.LoginRepository) UserUsecase {
	return &userUsecase{ur, lr}
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
func (u *userUsecase) PostUser(ctx *gin.Context) (*response.UserResponse, error) {

	// リクエストボディーの取得
	req := request.PostUserBodyRequest{}
	if err := req.Bind(ctx); err != nil {
		return nil, errors.New("faild request_body bind")
	}

	// 既存ユーザーがいればエラーを返す
	user, err := u.ur.GetUserInfo(req.UserID)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("This userID already used")
	}

	// ここにトランザクションを貼り、ログインデータ登録まで一連の処理にしたい
	err = infra.Transaction(func(tx *sqlx.Tx) error {

		// ユーザー登録
		err := u.ur.InsertUser(tx, entity.UserEntity{
			UserId:      req.UserID,
			MailAddress: *req.MailAddress,
		})
		if err != nil {
			return err
		}

		// ログイン情報
		err = u.lr.InsertLogin(tx, entity.LoginEntity{
			UserId:   req.UserID,
			Password: req.Password,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		User: response.User{
			UserId:      req.UserID,
			MailAddress: *req.MailAddress,
		},
		Login: response.Login{
			Password: req.Password,
		},
	}, nil
}
