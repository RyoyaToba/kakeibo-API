package mysql

import (
	"fmt"
	"strings"
	"your-project/api/entity"

	"github.com/jmoiron/sqlx"
)

type UserInformationMySQL interface {
	GetUserInfo(userID string) (*entity.UserInformation, error)
}

type userInformationMySQL struct {
	db *sqlx.DB
}

func NewUserInformationMySQL(db *sqlx.DB) UserInformationMySQL {
	return &userInformationMySQL{db: db}
}

func (u userInformationMySQL) GetUserInfo(userID string) (*entity.UserInformation, error) {
	// クエリを作成
	query := strings.Join([]string{
		"SELECT",
		"	user_id,",
		"	mail_address,",
		"	created_by,",
		"	created_at,",
		"	updated_by,",
		"	updated_at",
		"FROM",
		"	user_information",
		"WHERE",
		"	user_id = ? ",
	}, " ")

	var user entity.UserInformation
	err := u.db.Get(&user, query, userID)
	if err != nil {
		return nil, fmt.Errorf("データ取得エラー: %w", err)
	}

	return &user, nil
}
