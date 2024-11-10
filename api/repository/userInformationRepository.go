package repository

import (
	"database/sql"
	"errors"
	"your-project/api/entity"
)

type UserInformationRepositoryInt interface {
	GetUserInfo(userId string) (entity.UserInformation, error)
}

type UserInformationRepository struct {
	db *sql.DB
}

func NewUserInformationRepository(db *sql.DB) *UserInformationRepository {
	return &UserInformationRepository{db: db}
}

func (repo *UserInformationRepository) GetUserInfo(userId string) (entity.UserInformation, error) {
	var userInfo entity.UserInformation

	err := repo.db.QueryRow(
		"SELECT user_id, serial_number, mail_address, created_by, created_date, updated_by, updated_date FROM user_information WHERE user_id = $1", userId).
		Scan(&userInfo.UserId, &userInfo.SerialNumber, &userInfo.MailAddress, &userInfo.CreatedBy, &userInfo.CreatedDate, &userInfo.UpdatedBy, &userInfo.UpdatedDate)

	// エラーチェック
	if err != nil {
		// データが見つからなかった場合の処理
		if errors.Is(err, sql.ErrNoRows) {
			return entity.UserInformation{}, nil // 見つからなかった場合は空の構造体を返す
		}
		// 他のエラーの場合はそのまま返す
		return entity.UserInformation{}, err
	}

	return userInfo, nil
}
