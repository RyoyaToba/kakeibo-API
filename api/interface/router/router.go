package router

import (
	"your-project/api/interface/handler"

	"github.com/gin-gonic/gin"
)

// Handlers 構造体に、各エンドポイントのハンドラーをまとめる
type Handlers struct {
	UserInformationHandler handler.UserInformationHandler
	MessageHandler         handler.MessageHandler
}

// SetRouter はルーターを初期化し、エンドポイントを設定します。
func SetRouter(handlers *Handlers) *gin.Engine {
	r := gin.Default()

	// ヘルスチェック
	r.GET("/v1/healthCheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"check": "ok",
		})
	})

	// APIエンドポイント
	v1Routes := r.Group("/v1")
	{
		v1Routes.GET("/getUserInformation/:userId", handlers.UserInformationHandler.GetUserInfo)
		//v1Routes.GET("message/:id", handlers.MessageHandler.GetMessage) // 未使用
	}

	return r
}
