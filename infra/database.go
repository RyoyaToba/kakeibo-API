package infra

import (
	"database/sql"
	"fmt"
	env "your-project/api"
)

// SetupDB はDBのセットアップと接続を行い、*sql.DBを返します。
func SetupDB() (*sql.DB, error) {
	// DBConfigを取得
	dbConfig := env.LoadDBConfig()

	// PostgreSQL用のDSNを構築
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	// PostgreSQLに接続
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Database connection error: %v", err)
	}

	return db, nil
}
