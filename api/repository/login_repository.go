package repository

import (
	"your-project/api/entity"
	"your-project/api/mysql"
	"your-project/infra"

	"github.com/jmoiron/sqlx"
)

type LoginRepository interface {
	GetLoginInfo(userID string) (*entity.LoginEntity, error)
	InsertLogin(tx *sqlx.Tx, user entity.LoginEntity) error
}

type loginRepository struct {
}

func NewLoginRepository() LoginRepository {
	return &loginRepository{}
}

func (r *loginRepository) GetLoginInfo(userID string) (*entity.LoginEntity, error) {
	db := infra.SetupDB()
	li, err := mysql.NewLoginMySQL(db).GetLoginInfo(userID)
	if err != nil {
		return nil, err
	}
	return li, nil
}

func (r *loginRepository) InsertLogin(tx *sqlx.Tx, user entity.LoginEntity) error {
	db := infra.SetupDB()
	err := mysql.NewLoginMySQL(db).InsertLogin(tx, user)
	if err != nil {
		return err
	}
	return nil
}
