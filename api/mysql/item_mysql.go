package mysql

import (
	"fmt"
	"strings"
	"your-project/api/entity"

	"github.com/jmoiron/sqlx"
)

type ItemMySQL interface {
	GetItem(userID string, from string, to string) ([]*entity.Item, error)
	Insert(tx *sqlx.Tx, userID string, item entity.Item) error
	Update(tx *sqlx.Tx, userID string, item entity.Item) error
}

type itemMySQL struct {
	db *sqlx.DB
}

func NewItemMySQL(db *sqlx.DB) ItemMySQL {
	return &itemMySQL{db: db}
}

// GetItem .
func (im itemMySQL) GetItem(userID string, from string, to string) ([]*entity.Item, error) {
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
		"	AND target_date >= ? ",
		"	AND target_date <= ? ",
	}, " ")

	var items []*entity.Item
	err := im.db.Select(&items, query, userID, from, to)
	if err != nil {
		return nil, fmt.Errorf("データ取得エラー: %w", err)
	}

	return items, nil
}

// Insert .
func (im itemMySQL) Insert(tx *sqlx.Tx, userID string, item entity.Item) error {
	query := strings.Join([]string{
		"INSERT INTO `item`(",
		"	`user_id`, `name`, `price`, `target_date`, `category_id`, `bank_select_id`",
		") VALUES (?, ?, ?, ?, ?, ?)",
	}, " ")

	stmp, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmp.Close()
	_, err = stmp.Exec(
		userID,
		item.Name,
		item.Price,
		item.TargetDate,
		item.CategoryId,
		item.BankSelectId,
	)
	if err != nil {
		return err
	}
	return nil
}

// Update .
func (im itemMySQL) Update(tx *sqlx.Tx, userID string, item entity.Item) error {
	query := strings.Join([]string{
		"Update item SET (",
		"	`name` = ? ",
		"	`price` = ? ",
		"	`target_date` = ? ",
		"	`category_id` = ? ",
		"	`bank_select_id` = ? ",
		"	`updated_at` = NOW()",
		"WHERE",
		"	user_id = ? ",
		"	AND item_id = ? ",
	}, " ")
	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		item.Name,
		item.Price,
		item.TargetDate,
		item.CategoryId,
		item.BankSelectId,
		userID,
		item.ItemId,
	)
	if err != nil {
		return err
	}
	return nil
}
