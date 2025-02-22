package application

import (
	"your-project/api/repository"
	"your-project/api/response"
)

type ItemUsecase interface {
	Get(userID string) (*response.Item, error)
}

type itemUsecase struct {
	ir repository.ItemRepository
}

func NewItemUsecase(ir repository.ItemRepository) ItemUsecase {
	return itemUsecase{ir}
}

func (iu itemUsecase) Get(userID string) (*response.Item, error) {

	return &response.Item{}, nil
}
