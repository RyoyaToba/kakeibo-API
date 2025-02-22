package repository

import (
	"your-project/api/entity"
	"your-project/api/mysql"
	"your-project/infra"
)

type ItemRepository interface {
	GetItem(userID string, from string, to string) ([]*entity.Item, error)
}

type itemRepository struct {
}

func NewItemRepository() ItemRepository {
	return &itemRepository{}
}

func (ir *itemRepository) GetItem(userID string, from string, to string) ([]*entity.Item, error) {
	db := infra.SetupDB()
	items, err := mysql.NewItemMySQL(db).GetItem(userID, from, to)
	if err != nil {
		return nil, nil
	}
	return items, nil
}
