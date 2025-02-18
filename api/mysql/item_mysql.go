package mysql

import (
	"fmt"
	"strings"
	"your-project/api/entity"

	"github.com/jmoiron/sqlx"
)

type ItemMySQL interface {
	GetItem(userID string) ([]*entity.Item, error)
}

type itemMySQL struct {
	db *sqlx.DB
}

func NewItemMySQL(db *sqlx.DB) ItemMySQL {
	return &itemMySQL{db: db}
}

// GetItem .
func (im itemMySQL) GetItem(userID string) ([]*entity.Item, error) {
	// クエリを作成
	query := strings.Join([]string{
		"SELECT",
		"	item_id,",
		"	user_id,",
		"	name,",
		"	price,",
		"	target_date,",
		"	category_id,",
		"	bank_select_id,",
		"	created_by,",
		"	created_at,",
		"	updated_by,",
		"	updated_at",
		"FROM",
		"	item",
		"WHERE",
		"	user_id = ? ",
	}, " ")

	var items []*entity.Item
	err := im.db.Select(&items, query, userID)
	if err != nil {
		return nil, fmt.Errorf("データ取得エラー: %w", err)
	}

	return items, nil
}
