package main

import (
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

	// ルーターをセットアップ
	r := router.SetRouter(handlers)

	// サーバー起動
	r.Run(":8080") // Ginの `Run()` を使用
}
