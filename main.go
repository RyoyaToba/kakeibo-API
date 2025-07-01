package main

import (
	usecase "your-project/api/application"
	"your-project/api/interface/handler"
	"your-project/api/interface/router"
	"your-project/api/repository"
)

func main() {
	uu := usecase.NewUserUsecase(repository.NewUserRepository(), repository.NewLoginRepository())
	iu := usecase.NewItemUsecase(repository.NewItemRepository())
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
