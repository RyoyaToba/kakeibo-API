package main

import (
	"database/sql"
	"fmt"
	"net/http"
	env "your-project/api"
	usecase "your-project/api/application"
	"your-project/api/interface/handler"
	"your-project/api/repository"
	"your-project/api/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	// DBConfigを取得
	dbConfig := env.LoadDBConfig()

	// MySQLデータベースのDSNを構築
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	// postgres用のDSNを構築
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	// postgresに接続
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Sprintf("Database connection error: %v", err))
	}
	defer db.Close()

	// RepositoryとUsecaseの初期化
	//messageRepo := repository.NewMySQLMessageRepository(db)
	//messageUsecase := usecase.NewMessageUsecase(messageRepo)

	userInformationRepo := repository.NewUserInformationRepository(db)
	userInformationService := service.NewUserInformationService(userInformationRepo)
	userInfomationUsecase := usecase.NewUserInformationUsecase(userInformationService)

	// Handlerの初期化
	//messageHandler := handler.NewMessageHandler(messageUsecase)
	userInformationHandler := handler.NewUserInformationHandler(userInfomationUsecase)

	// Ginのルーターを初期化
	router := gin.New()

	// ルートを設定
	router.GET("/v1/helthCheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"check": "ok",
		})
	})
	//router.GET("/v1/getmessage/:id", messageHandler.GetMessage)

	router.GET("/v1/getUserInfomation/:userId", userInformationHandler.GetUserInfo)

	// サーバーを起動
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	httpServer.ListenAndServe()
}
