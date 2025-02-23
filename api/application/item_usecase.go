package application

import (
	"your-project/api/repository"
	"your-project/api/response"

	"github.com/gin-gonic/gin"
)

type ItemUsecase interface {
	Get(ctx *gin.Context) (*response.ItemResponse, error)
	Post(ctx *gin.Context) (*response.ItemResponse, error)
	Put(ctx *gin.Context) (*response.ItemResponse, error)
}

type itemUsecase struct {
	ir repository.ItemRepository
}

func NewItemUsecase(ir repository.ItemRepository) ItemUsecase {
	return itemUsecase{ir}
}

// GET /v1/item
func (iu itemUsecase) Get(ctx *gin.Context) (*response.ItemResponse, error) {
	// パラメータの取得
	from := ctx.Query("from")
	to := ctx.Query("to")
	userID := ctx.Query("userID")

	// アイテム情報を取得
	items, err := iu.ir.GetItem(userID, from, to)
	if err != nil {
		return nil, err
	}

	// レスポンス作成
	itemList := make([]*response.Item, len(items))
	for i, item := range items {
		itemList[i] = &response.Item{
			UserID:       userID,
			ItemId:       item.ItemId,
			Name:         item.Name,
			Price:        item.Price,
			TargetDate:   item.TargetDate,
			CategoryId:   item.CategoryId,
			BankSelectId: item.BankSelectId,
		}
	}
	return &response.ItemResponse{
		Items: itemList,
	}, nil
}

// POST /v1/item
func (iu itemUsecase) Post(ctx *gin.Context) (*response.ItemResponse, error) {
	return nil, nil
}

// PUT /v1/item
func (iu itemUsecase) Put(ctx *gin.Context) (*response.ItemResponse, error) {
	return nil, nil
}
