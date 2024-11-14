// handler/message.go

package handler

import (
	"net/http"
	usecase "your-project/api/application"

	"github.com/gin-gonic/gin"
)

// MessageHandler はメッセージエンドポイントのハンドラ
type MessageHandler struct {
	usecase usecase.MessageUsecase
}

// NewMessageHandler は MessageHandler を初期化するコンストラクタ
func NewMessageHandler(uc usecase.MessageUsecase) *MessageHandler {
	return &MessageHandler{usecase: uc}
}

// GetMessage は指定された ID のメッセージを取得するエンドポイント
func (h *MessageHandler) GetMessage(c *gin.Context) {
	id := c.Param("id")
	message, err := h.usecase.GetMessage(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}
