package models

// Budget represents a budget for a user, containing a list of accounts.
type Budget struct {
	ID       int               `json:"id"` // Unique identifier for the budget.
	UserID   int               `json:"user_id"` // The ID of the user this budget belongs to.
	Name     string            `json:"name"` // The name of the budget.
	Amount   float64           `json:"amount"` // The total amount allocated to this budget.
	Accounts []Account `json:"accounts"` // A list of accounts associated with this budget.
}
