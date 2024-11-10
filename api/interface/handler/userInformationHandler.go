// handler/message.go

package handler

import (
	"log"
	"net/http"
	usecase "your-project/api/application"

	"github.com/gin-gonic/gin"
)

// MessageHandler はメッセージエンドポイントのハンドラ
type UserInformationHandler struct {
	usecase usecase.UserInformationUsecase
}

// NewMessageHandler は MessageHandler を初期化するコンストラクタ
func NewUserInformationHandler(uc usecase.UserInformationUsecase) *UserInformationHandler {
	return &UserInformationHandler{usecase: uc}
}

// GetMessage は指定された ID のメッセージを取得するエンドポイント
func (h *UserInformationHandler) GetUserInfo(c *gin.Context) {
	userId := c.Param("userId")
	log.Println("userId = ", userId)
	userInformation, err := h.usecase.GetUserInfo(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"userInfomation": userInformation})
}
