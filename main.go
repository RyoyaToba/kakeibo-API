package main

import (
	"net/http"
	usecase "your-project/api/application"
	"your-project/api/interface/handler"
	"your-project/api/interface/router"
	"your-project/api/repository"
	"your-project/api/service"
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
	ur := repository.NewUserInformationRepository(db)
	us := service.NewUserInformationService(ur)
	uc := usecase.NewUserInformationUsecase(us)
	userInformationHandler := handler.NewUserInformationHandler(uc)

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
