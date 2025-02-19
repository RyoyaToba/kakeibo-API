package mysql

import (
	"os"
	"testing"
	"your-project/api/entity"
	"your-project/infra"

	"github.com/stretchr/testify/assert"
)

func Test_GetItem(t *testing.T) {
	os.Setenv("GO_TEST_ENV", "true")
	defer os.Setenv("GO_TEST_ENV", "false")
	db := infra.SetupDB()
	defer db.Close()

	tests := map[string]struct {
		input  func() (string, string, string)
		output func(*testing.T, []*entity.Item, error)
	}{
		"success_some_items": {
			input: func() (string, string, string) {
				db.Exec(`INSERT INTO item (user_id, name, price, target_date, category_id, bank_select_id, created_by, created_at, updated_by, updated_at) 
							VALUES ('test1', '家賃', 100000, '2025-02-18', 1, 1, 'admin', NOW(), 'admin', NOW())`)
				return "test1", "2025-02-01", "2025-02-28"
			},
			output: func(t *testing.T, items []*entity.Item, err error) {
				assert.Equal(t, 1, len(items))
			},
		},
	}

	e := NewItemMySQL(db)
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			userID, from, to := tt.input()
			items, err := e.GetItem(userID, from, to)
			tt.output(t, items, err)
		})
	}
	db.Exec("TRUNCATE TABLE item")
}
