package main

import (
	"log"
	usecase "your-project/api/application"
	"your-project/api/interface/handler"
	"your-project/api/interface/router"
	"your-project/api/repository"
	"your-project/api/service"
)

func main() {
	ur := repository.NewUserInformationRepository()
	us := service.NewUserInformationService(ur)
	uu := usecase.NewUserInformationUsecase(us)
	ir := repository.NewItemRepository()
	iu := usecase.NewItemUsecase(ir)
	// ハンドラーをセットアップ
	handlers := &router.Handlers{
		UserInformationHandler: handler.NewUserInformationHandler(uu),
		ItemHandler:            handler.NewItemHandler(iu),
	}

	// デバッグ用ログ
	if handlers == nil {
		log.Fatal("handlers is nil!")
	}
	if handlers.UserInformationHandler == nil {
		log.Fatal("handlers.UserInformationHandler is nil!")
	}
	if handlers.ItemHandler == nil {
		log.Fatal("handlers.ItemHandler is nil!")
	}

	// ルーターをセットアップ
	r := router.SetRouter(handlers)

	// サーバー起動
	r.Run(":8080") // Ginの `Run()` を使用
}
