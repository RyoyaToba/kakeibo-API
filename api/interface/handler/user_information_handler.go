// handler/message.go

package handler

import (
	"net/http"
	usecase "your-project/api/application"

	"github.com/gin-gonic/gin"
)

type UserInformationHandler interface {
	GetUserInfo(ctx *gin.Context)
}

type userInformationHandler struct {
	usecase usecase.UserInformationUsecase
}

func NewUserInformationHandler(uc usecase.UserInformationUsecase) UserInformationHandler {
	return userInformationHandler{uc}
}

func (uc userInformationHandler) GetUserInfo(c *gin.Context) {
	userId := c.Param("userId")
	userInformation, err := uc.usecase.GetUserInfo(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"userInfomation": userInformation})
}
