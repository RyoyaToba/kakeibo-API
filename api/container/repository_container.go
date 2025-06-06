package container

import "your-project/api/repository"

// RepositoryContainer は全てのリポジトリを管理
type RepositoryContainer struct {
	UserInformationRepository repository.UserInformationRepository
	ItemRepository            repository.ItemRepository
}

// NewRepositoryContainer は Repository をまとめて初期化する
func NewRepositoryContainer() *RepositoryContainer {
	return &RepositoryContainer{
		UserInformationRepository: repository.NewUserInformationRepository(),
		ItemRepository:            repository.NewItemRepository(),
	}
}
