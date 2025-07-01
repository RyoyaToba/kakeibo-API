package container

import "your-project/api/repository"

// RepositoryContainer は全てのリポジトリを管理
type RepositoryContainer struct {
	UserRepository  repository.UserRepository
	LoginRepository repository.LoginRepository
	ItemRepository  repository.ItemRepository
}

// NewRepositoryContainer は Repository をまとめて初期化する
func NewRepositoryContainer() *RepositoryContainer {
	return &RepositoryContainer{
		UserRepository:  repository.NewUserRepository(),
		LoginRepository: repository.NewLoginRepository(),
		ItemRepository:  repository.NewItemRepository(),
	}
}
