package main

import (
	usecase "your-project/api/application"
	"your-project/api/interface/handler"
	"your-project/api/interface/router"
	"your-project/api/repository"
)

func main() {
	ur := repository.NewUserRepository()
	uu := usecase.NewUserUsecase(ur)
	ir := repository.NewItemRepository()
	iu := usecase.NewItemUsecase(ir)
	// ハンドラーをセットアップ
	handlers := &router.Handlers{
		UserHandler: handler.NewUserHandler(uu),
		ItemHandler: handler.NewItemHandler(iu),
	}

	// ルーターをセットアップ
	r := router.SetRouter(handlers)

	// サーバー起動
	r.Run(":8080") // Ginの `Run()` を使用
}
