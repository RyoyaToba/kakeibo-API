package entity

import "time"

type LoginEntity struct {
	UserId    string    `db:"user_id"`
	Password  string    `db:"password"`
	CreatedBy string    `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedBy string    `db:"updated_by"`
	UpdatedAt time.Time `db:"updated_at"`
}
