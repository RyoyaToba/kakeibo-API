package entity

import (
	"time"
)

type Item struct {
	ItemId       int64     `db:"item_id"`
	UserId       string    `db:"user_id"`
	Name         string    `db:"name"`
	Price        int64     `db:"price"`
	TargetDate   string    `db:"target_date"`
	CategoryId   int64     `db:"category_id"`
	BankSelectId int64     `db:"bank_select_id"`
	CreatedBy    string    `db:"created_by"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedBy    string    `db:"updated_by"`
	UpdatedAt    time.Time `db:"updated_at"`
}
