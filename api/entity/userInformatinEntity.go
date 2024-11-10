package entity

import (
	"database/sql"
	"time"
)

type UserInformation struct {
	UserId       string
	SerialNumber int8
	MailAddress  sql.NullString
	CreatedBy    string
	CreatedDate  time.Time
	UpdatedBy    string
	UpdatedDate  time.Time
}
