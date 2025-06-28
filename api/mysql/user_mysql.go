package mysql

import (
	"fmt"
	"strings"
	"your-project/api/entity"

	"github.com/jmoiron/sqlx"
)

type UserMySQL interface {
	GetUserInfo(userID string) (*entity.UserEntity, error)
	InsertUser(user entity.UserEntity) error
}

type userMySQL struct {
	db *sqlx.DB
}

func NewUserMySQL(db *sqlx.DB) UserMySQL {
	return &userMySQL{db: db}
}

// GetUserInfo .
func (u userMySQL) GetUserInfo(userID string) (*entity.UserEntity, error) {
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
		"	user",
		"WHERE",
		"	user_id = ? ",
	}, " ")

	var user entity.UserEntity
	err := u.db.Get(&user, query, userID)
	if err != nil {
		return nil, fmt.Errorf("データ取得エラー: %w", err)
	}

	return &user, nil
}

// InsertUser .
func (u userMySQL) InsertUser(user entity.UserEntity) error {
	// クエリを作成
	query := strings.Join([]string{
		"INSERT INTO `user`(`user_id`, `mail_address`, `created_by`, `updated_by`)",
		"VALUES",
		"(?, ?, ?, ?)",
	}, " ")

	stmp, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmp.Close()
	_, err = stmp.Exec(
		user.UserId,
		user.MailAddress,
		user.UserId,
		user.UserId,
	)
	if err != nil {
		return err
	}
	return nil
}
