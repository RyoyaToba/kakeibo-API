package container

import "your-project/api/repository"

// RepositoryContainer は全てのリポジトリを管理
type RepositoryContainer struct {
	UserRepository repository.UserRepository
	ItemRepository repository.ItemRepository
}

// NewRepositoryContainer は Repository をまとめて初期化する
func NewRepositoryContainer() *RepositoryContainer {
	return &RepositoryContainer{
		UserRepository: repository.NewUserRepository(),
		ItemRepository: repository.NewItemRepository(),
	}
}
