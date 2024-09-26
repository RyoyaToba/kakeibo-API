// main.go

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	usecase "your-project/api/application"
	"your-project/api/interface/handler"
	"your-project/api/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// MySQLデータベースの接続情報
const (
	dsn = "root:rootpassword@tcp(mysql:3306)/testdb"
)

func main() {
	// MySQLに接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Database connection error: %v", err))
	}
	defer db.Close()

	// RepositoryとUsecaseの初期化
	messageRepo := repository.NewMySQLMessageRepository(db)
	messageUsecase := usecase.NewMessageUsecase(messageRepo)

	// Handlerの初期化
	messageHandler := handler.NewMessageHandler(messageUsecase)

	// Ginのルーターを初期化
	router := gin.Default()

	// ルートを設定
	router.GET("/v1/getmessage/:id", messageHandler.GetMessage)

	// サーバーを起動
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	httpServer.ListenAndServe()
}
