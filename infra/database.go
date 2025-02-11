package infra

import (
	"fmt"
	"log"
	env "your-project/api"

	_ "github.com/go-sql-driver/mysql" // MySQL ドライバをインポート
	"github.com/jmoiron/sqlx"
)

// SetupDB DBのセットアップと接続を行い、*sqlx.DBを返します。
func SetupDB() *sqlx.DB {
	// 環境変数からDBConfigを取得
	dbConfig := env.LoadDBConfig()

	// MySQL用のDSNを構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name,
	)

	log.Printf("Connecting to DB with DSN: %s", dsn) // DSNをログに出力

	// MySQLに接続
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("DB接続エラー: %v", err)
	}

	return db
}
