package container

import "your-project/api/interface/handler"

// HandlerContainer は全ての Handler を管理
type HandlerContainer struct {
	UserHandler handler.UserHandler
	ItemHandler handler.ItemHandler
}

// NewHandlerContainer は UsecaseContainer を受け取って Handler を初期化
func NewHandlerContainer(usecase *UsecaseContainer) *HandlerContainer {
	return &HandlerContainer{
		UserHandler: handler.NewUserHandler(usecase.UserUsecase),
		ItemHandler: handler.NewItemHandler(usecase.ItemUsecase),
	}
}
