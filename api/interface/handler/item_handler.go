package handler

import (
	"net/http"
	usecase "your-project/api/application"

	"github.com/gin-gonic/gin"
)

type ItemHandler interface {
	Get(*gin.Context)
}

type itemHandler struct {
	usecase usecase.ItemUsecase
}

func NewItemHandler(iu usecase.ItemUsecase) ItemHandler {
	return itemHandler{iu}
}

// GET
func (ih itemHandler) Get(ctx *gin.Context) {
	item, err := ih.usecase.Get(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"item": item})
}
