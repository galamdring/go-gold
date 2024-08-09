package models

import "github.com/shopspring/decimal"

type Transaction struct {
	ID         int             `json:"id"`
	AccountID  int             `json:"account_id"`
	Amount     decimal.Decimal `json:"amount"`
	CategoryID int             `json:"category_id"`
	Note       string          `json:"note"`
}
