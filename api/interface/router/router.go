package router

import (
	"time"
	"your-project/api/interface/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Handlers 構造体に、各エンドポイントのハンドラーをまとめる
type Handlers struct {
	UserHandler    handler.UserHandler
	MessageHandler handler.MessageHandler
	ItemHandler    handler.ItemHandler
}

// SetRouter はルーターを初期化し、エンドポイントを設定します。
func SetRouter(handlers *Handlers) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // フロントエンドのURL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ヘルスチェック
	r.GET("/v1/healthCheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"check": "ok",
		})
	})

	// APIエンドポイント
	v1Routes := r.Group("/v1")
	{
		// GET user
		v1Routes.GET("/getUserInformation/:userId", handlers.UserHandler.GetUserInfo)
		// GET item
		v1Routes.GET("/item", handlers.ItemHandler.Get)
	}

	return r
}
