package repository

import "expense-tracker/internal/domain"

type ExpenseRepository interface {
	Create(exp domain.Expense) error
	Get(category, from, to, sort string) ([]domain.Expense, float64, error)
	GetSummary(userID string) (map[string]float64, error)
}
