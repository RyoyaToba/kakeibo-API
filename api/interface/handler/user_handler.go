// handler/message.go

package handler

import (
	"net/http"
	usecase "your-project/api/application"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUserInfo(ctx *gin.Context)
}

type userHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) UserHandler {
	return userHandler{uc}
}

func (uc userHandler) GetUserInfo(c *gin.Context) {
	userId := c.Param("userId")
	user, err := uc.usecase.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
