package mysql

import (
	"os"
	"testing"
	"your-project/api/entity"

	"your-project/infra"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func Test_GetUserInfo(t *testing.T) {
	os.Setenv("GO_TEST_ENV", "true")
	defer os.Setenv("GO_TEST_ENV", "false")
	db := infra.SetupDB()
	defer db.Close()

	tests := map[string]struct {
		input  func() string
		output func(*testing.T, *entity.UserInformation, error)
	}{
		"success": {
			input: func() string {
				db.Exec(`INSERT INTO user_information (user_id, mail_address, created_by, created_at, updated_by, updated_at)
							VALUES ('test1', 'test@example.com', 'admin', NOW(), 'admin', NOW())`)
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
