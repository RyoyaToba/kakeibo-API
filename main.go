package main

import (
	"your-project/api/application"
	"your-project/api/interface/handler"
	"your-project/api/interface/router"
	"your-project/api/repository"
	"your-project/api/service"
)

func main() {
	ur := repository.NewUserInformationRepository()
	us := service.NewUserInformationService(ur)
	uc := application.NewUserInformationUsecase(us)
	// ハンドラーをセットアップ
	handlers := &router.Handlers{
		UserInformationHandler: handler.NewUserInformationHandler(uc),
		// MessageHandler: handler.NewMessageHandler(), // 必要なら追加
	}

	// ルーターをセットアップ
	r := router.SetRouter(handlers)

	// サーバー起動
	r.Run(":8080") // Ginの `Run()` を使用
}
