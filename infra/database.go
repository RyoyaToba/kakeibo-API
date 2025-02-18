package infra

import (
	"fmt"
	"log"
	"os"
	env "your-project/api"

	_ "github.com/go-sql-driver/mysql" // MySQL ドライバをインポート
	"github.com/jmoiron/sqlx"
)

func SetupDB() *sqlx.DB {
	// 環境変数からDBConfigを取得
	dbConfig := env.LoadDBConfig()

	// テスト実行中なら、テスト用のDBホストを適用
	if isTestEnv() {
		dbConfig.User = "testuser"
		dbConfig.Password = "testpass"
		dbConfig.Host = "127.0.0.1"
		dbConfig.Port = "3307"
		dbConfig.Name = "testdb"
	}

	// MySQL用のDSNを構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name,
	)

	log.Printf("Connecting to DB with DSN: %s", dsn)

	// MySQLに接続
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("DB接続エラー: %v", err)
	}

	return db
}

// 環境変数 `GO_TEST_ENV` がセットされているか確認
func isTestEnv() bool {
	return os.Getenv("GO_TEST_ENV") == "true"
}
