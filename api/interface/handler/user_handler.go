// handler/message.go

package handler

import (
	"net/http"
	usecase "your-project/api/application"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUserInfo(ctx *gin.Context)
	PostUser(ctx *gin.Context)
}

type userHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) UserHandler {
	return userHandler{uc}
}

func (u userHandler) GetUserInfo(c *gin.Context) {
	userId := c.Param("userId")
	user, err := u.usecase.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u userHandler) PostUser(ctx *gin.Context) {
	res, err := u.usecase.PostUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"res": res})
}
