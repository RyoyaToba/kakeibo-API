package container

// MainContainer はアプリ全体の依存関係を統括
type MainContainer struct {
	Repository *RepositoryContainer
	Usecase    *UsecaseContainer
	Handler    *HandlerContainer
}

// NewMainContainer はすべてのコンテナを統合
func NewMainContainer() *MainContainer {
	repo := NewRepositoryContainer()
	usecase := NewUsecaseContainer(repo)
	handler := NewHandlerContainer(usecase)

	return &MainContainer{
		Repository: repo,
		Usecase:    usecase,
		Handler:    handler,
	}
}
