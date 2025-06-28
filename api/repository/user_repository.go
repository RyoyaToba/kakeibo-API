package repository

import (
	"your-project/api/entity"
	"your-project/api/mysql"
	"your-project/infra"
)

type UserRepository interface {
	GetUserInfo(userID string) (entity.UserEntity, error)
	InsertUser(user entity.UserEntity) error
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// GetUserInfo ユーザー情報を取得します
func (ur *userRepository) GetUserInfo(userID string) (entity.UserEntity, error) {
	db := infra.SetupDB()
	user, err := mysql.NewUserMySQL(db).GetUserInfo(userID)
	if err != nil {
		return entity.UserEntity{}, nil
	}
	return *user, nil
}

// InsertUser ユーザー登録
func (ur *userRepository) InsertUser(user entity.UserEntity) error {
	db := infra.SetupDB()
	return mysql.NewUserMySQL(db).InsertUser(user)
}
