package response

type ItemResponse struct {
	Items []*Item `json:"items"`
}

type Item struct {
	UserID       string `json:"user_id"`
	ItemId       int64  `json:"item_id"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	TargetDate   string `json:"target_date"`
	CategoryId   int64  `json:"category_id"`
	BankSelectId int64  `json:"bank_select_id"`
}
