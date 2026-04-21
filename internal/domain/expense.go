package domain

type Expense struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	ExpenseDate string  `json:"expense_date"`
	CreatedAt   string  `json:"created_at"`
}
