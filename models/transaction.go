package models

import "github.com/shopspring/decimal"

type Transaction struct {
	ID         int             `json:"id"`
	AccountID  int             `json:"accountId"`
	Amount     decimal.Decimal `json:"amount"`
	CategoryID int             `json:"categoryId"`
	Note       string          `json:"note"`
}
