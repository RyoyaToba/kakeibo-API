package entity

import (
	"time"
)

type UserEntity struct {
	UserId      string    `db:"user_id"`
	MailAddress string    `db:"mail_address"`
	CreatedBy   string    `db:"created_by"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedBy   string    `db:"updated_by"`
	UpdatedAt   time.Time `db:"updated_at"`
}
