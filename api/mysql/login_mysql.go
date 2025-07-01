package mysql

import (
	"strings"
	"your-project/api/entity"

	"github.com/jmoiron/sqlx"
)

type LoginMySQL interface {
	GetLoginInfo(userID string) (*entity.LoginEntity, error)
	InsertLogin(tx *sqlx.Tx, user entity.LoginEntity) error
}

type loginMySQL struct {
	db *sqlx.DB
}

func NewLoginMySQL(db *sqlx.DB) LoginMySQL {
	return &loginMySQL{db: db}
}

// GetLoginInfo .
func (m *loginMySQL) GetLoginInfo(userID string) (*entity.LoginEntity, error) {
	query := strings.Join([]string{
		"SELECT",
		"	`user_id`,",
		"	`password`,",
		"FROM",
		"	`login_info_master`",
		"WHERE",
		"	`user_id` = ?",
	}, " ")
	var li entity.LoginEntity
	err := m.db.Select(&li, query, userID)
	if err != nil {
		return nil, err
	}
	return &li, nil
}

// InsertLogin .
func (m *loginMySQL) InsertLogin(tx *sqlx.Tx, li entity.LoginEntity) error {
	query := strings.Join([]string{
		"INSERT INTO `login_info_master` (`user_id`, `password`, `created_by`, `updated_by`)",
		"VALUES",
		"(?, ?, ?, ?)",
	}, " ")

	stmp, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmp.Close()
	_, err = stmp.Exec(
		li.UserId,
		li.Password,
		li.UserId,
		li.UserId,
	)
	if err != nil {
		return err
	}
	return nil
}
