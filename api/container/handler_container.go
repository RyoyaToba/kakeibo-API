package container

import "your-project/api/interface/handler"

// HandlerContainer は全ての Handler を管理
type HandlerContainer struct {
	UserInformationHandler handler.UserInformationHandler
	ItemHandler            handler.ItemHandler
}

// NewHandlerContainer は UsecaseContainer を受け取って Handler を初期化
func NewHandlerContainer(usecase *UsecaseContainer) *HandlerContainer {
	return &HandlerContainer{
		UserInformationHandler: handler.NewUserInformationHandler(usecase.UserInformationUsecase),
		ItemHandler:            handler.NewItemHandler(usecase.ItemUsecase),
	}
}
