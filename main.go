package main

import (
	"net/http"
	"your-project/api/interface/handler"
	"your-project/api/interface/router"
	"your-project/infra"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {

	// DBのセットアップ
	db, err := infra.SetupDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//messageHandler := handler.NewMessageHandler(messageUsecase)
	userInformationHandler := handler.NewUserInformationHandler(db)

	// Handlers構造体にハンドラーをまとめる
	handlers := &router.Handlers{
		UserInformationHandler: userInformationHandler,
		//MessageHandler:         messageHandler,
	}

	// Ginのルーターを初期化
	router := router.SetRouter(handlers)

	// サーバーを起動
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	httpServer.ListenAndServe()
}
