package request

import (
	"log"

	"github.com/gin-gonic/gin"
)

type PostUserBodyRequest struct {
	UserID      string  `json:"user_id"`
	MailAddress *string `json:"mail_address"`
}

func (r *PostUserBodyRequest) Bind(ctx *gin.Context) error {
	err := ctx.ShouldBindJSON(r)
	if err != nil {
		log.Printf("JSON bind error: %v", err)
		return err
	}
	return nil
}
