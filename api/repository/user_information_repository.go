package repository

import (
	"your-project/api/entity"
	"your-project/api/mysql"
	"your-project/infra"
)

type UserInformationRepository interface {
	GetUserInfo(userID string) (entity.UserInformation, error)
}

type userInformationRepository struct {
}

func NewUserInformationRepository() UserInformationRepository {
	return &userInformationRepository{}
}

// GetUserInfo ユーザー情報を取得します
func (ur *userInformationRepository) GetUserInfo(userID string) (entity.UserInformation, error) {

	db := infra.SetupDB()
	userInformation, err := mysql.NewUserInformationMySQL(db).GetUserInfo(userID)
	if err != nil {
		return entity.UserInformation{}, nil
	}
	return *userInformation, nil
}
