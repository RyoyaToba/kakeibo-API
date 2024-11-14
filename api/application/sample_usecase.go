package application

import (
	"errors"
	"your-project/api/repository"
)

// MessageUsecase はメッセージに関するビジネスロジックを定義するインターフェース
type MessageUsecase interface {
	GetMessage(id string) (string, error)
}

// MessageUsecaseImpl は MessageUsecase の具体的な実装
type MessageUsecaseImpl struct {
	repo repository.MessageRepository
}

// NewMessageUsecase は MessageUsecaseImpl を初期化するコンストラクタ
func NewMessageUsecase(repo repository.MessageRepository) *MessageUsecaseImpl {
	return &MessageUsecaseImpl{repo: repo}
}

// GetMessage は指定された ID のメッセージを取得する
func (uc *MessageUsecaseImpl) GetMessage(id string) (string, error) {
	message, err := uc.repo.GetMessageByID(id)
	if err != nil {
		return "", err
	}
	if message == "" {
		return "", errors.New("message not found")
	}
	return message, nil
}
