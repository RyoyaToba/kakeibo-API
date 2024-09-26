// repository/message.go

package repository

import (
	"database/sql"
	"errors"
)

// MessageRepository はメッセージに関するデータ操作のインターフェース
type MessageRepository interface {
	GetMessageByID(id string) (string, error)
}

// MySQLMessageRepository は MySQL を使った MessageRepository の実装
type MySQLMessageRepository struct {
	db *sql.DB
}

// NewMySQLMessageRepository は MySQLMessageRepository を初期化するコンストラクタ
func NewMySQLMessageRepository(db *sql.DB) *MySQLMessageRepository {
	return &MySQLMessageRepository{db: db}
}

// GetMessageByID は指定された ID に対応するメッセージを取得する
func (repo *MySQLMessageRepository) GetMessageByID(id string) (string, error) {
	var message string
	err := repo.db.QueryRow("SELECT message FROM test WHERE id = ?", id).Scan(&message)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil // 見つからなかった場合は空文字列を返す
		}
		return "", err
	}
	return message, nil
}
