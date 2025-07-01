package container

import usecase "your-project/api/application"

// UsecaseContainer は全ての Usecase を管理
type UsecaseContainer struct {
	UserUsecase usecase.UserUsecase
	ItemUsecase usecase.ItemUsecase
}

// NewUsecaseContainer は ServiceContainer と RepositoryContainer を受け取って Usecase を初期化
func NewUsecaseContainer(r *RepositoryContainer) *UsecaseContainer {
	return &UsecaseContainer{
		UserUsecase: usecase.NewUserUsecase(r.UserRepository, r.LoginRepository),
		ItemUsecase: usecase.NewItemUsecase(r.ItemRepository),
	}
}
