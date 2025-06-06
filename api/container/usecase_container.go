package container

import usecase "your-project/api/application"

// UsecaseContainer は全ての Usecase を管理
type UsecaseContainer struct {
	UserInformationUsecase usecase.UserInformationUsecase
	ItemUsecase            usecase.ItemUsecase
}

// NewUsecaseContainer は ServiceContainer と RepositoryContainer を受け取って Usecase を初期化
func NewUsecaseContainer(repo *RepositoryContainer) *UsecaseContainer {
	return &UsecaseContainer{
		UserInformationUsecase: usecase.NewUserInformationUsecase(repo.UserInformationRepository),
		ItemUsecase:            usecase.NewItemUsecase(repo.ItemRepository),
	}
}
