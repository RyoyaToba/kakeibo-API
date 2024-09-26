package main

import (
	"database/sql"
	"fmt"
	"net/http"
	env "your-project/api"
	usecase "your-project/api/application"

	"your-project/api/interface/handler"
	"your-project/api/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// DBConfigを取得
	dbConfig := env.LoadDBConfig()

	// MySQLデータベースのDSNを構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

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
