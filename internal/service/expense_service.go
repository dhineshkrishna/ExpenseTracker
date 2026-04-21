package service

import (
	"expense-tracker/internal/domain"
	"expense-tracker/internal/repository"
	"time"

	"github.com/google/uuid"
)

type ExpenseService struct {
	Repo repository.ExpenseRepository
}

func NewExpenseService(r repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{Repo: r}
}

func (s *ExpenseService) CreateExpense(e domain.Expense) (domain.Expense, error) {

	if e.ID == "" {
		e.ID = uuid.New().String()
	}

	e.CreatedAt = time.Now().Format(time.RFC3339)

	err := s.Repo.Create(e)
	return e, err
}

func (s *ExpenseService) GetExpenses( category, from, to, sort string) ([]domain.Expense, float64, error) {
	return s.Repo.Get(category, from, to, sort)
}

func (s *ExpenseService) GetSummary() (map[string]float64, error) {
	return s.Repo.GetSummary()
}
