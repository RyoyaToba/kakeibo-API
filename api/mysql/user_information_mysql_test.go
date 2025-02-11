package mysql

import (
	"fmt"
	"testing"
	"your-project/api/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

const testDNS = "testuser:testpass@tcp(127.0.0.1:3307)/testdb?parseTime=true"

func setupTestDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", testDNS)
	if err != nil {
		return nil, err
	}

	// 接続確認
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("DBへのPing失敗: %w", err)
	}
	return db, nil
}

func Test_GetUserInfo(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("DB接続エラー: %v", err)
	}
	db.Exec(`INSERT INTO user_information (user_id, mail_address, created_by, created_at, updated_by, updated_at)
				VALUES ('test1', 'test@example.com', 'admin', NOW(), 'admin', NOW())`)
	defer db.Close()

	tests := map[string]struct {
		input  func() string
		output func(*testing.T, *entity.UserInformation, error)
	}{
		"success": {
			input: func() string {
				return "test1"
			},
			output: func(t *testing.T, ui *entity.UserInformation, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "test1", ui.UserId)
				assert.Equal(t, "test@example.com", ui.MailAddress)
			},
		},
	}

	e := NewUserInformationMySQL(db)
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			userID := tt.input()
			userInfo, err := e.GetUserInfo(userID)
			tt.output(t, userInfo, err)
		})
	}
	db.Exec("TRUNCATE TABLE user_information")

}
